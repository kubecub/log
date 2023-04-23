// Copyright © 2023 KubeCub. All rights reserved.
//
// Licensed under the MIT License (the "License");
// you may not use this file except in compliance with the License.

// After calling the log print function, you also need to print these logs to the supported output,
// so you need to implement the write function, whose write logic is stored in the `entry.go` file.
package cuslog

import (
	"bytes"
	"runtime"
	"strings"
	"time"
)

// Type is used to store all log information, that is, log configuration and log content.
// The writing logic is done around an instance of the Entry type.
type Entry struct {
	logger *logger
	Buffer *bytes.Buffer
	Map    map[string]interface{}
	Level  Level
	Time   time.Time
	File   string
	Line   int
	Func   string
	Format string
	Args   []interface{}
}

func entry(logger *logger) *Entry {
	return &Entry{logger: logger, Buffer: new(bytes.Buffer), Map: make(map[string]interface{}, 5)}
}

func (e *Entry) write(level Level, format string, args ...interface{}) {
	if e.logger.opt.level > level {
		return
	}
	e.Time = time.Now()
	e.Level = level
	e.Format = format
	e.Args = args
	if !e.logger.opt.disableCaller {
		// runtime.Caller() gets the file name and line number.
		// When calling runtime.Caller(), be careful to pass in the correct stack depth.
		if pc, file, line, ok := runtime.Caller(2); !ok {
			e.File = "???"
			e.Func = "???"
		} else {
			e.File, e.Line, e.Func = file, line, runtime.FuncForPC(pc).Name()
			e.Func = e.Func[strings.LastIndex(e.Func, "/")+1:]
		}
	}
	e.format()
	e.writer()
	e.release()
}

func (e *Entry) format() {
	_ = e.logger.opt.formatter.Format(e)
}

func (e *Entry) writer() {
	e.logger.mu.Lock()
	_, _ = e.logger.opt.output.Write(e.Buffer.Bytes())
	e.logger.mu.Unlock()
}

func (e *Entry) release() {
	e.Args, e.Line, e.File, e.Format, e.Func = nil, 0, "", "", ""
	e.Buffer.Reset()
	e.logger.entryPool.Put(e)
}
