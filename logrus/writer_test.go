// Copyright © 2026 KubeCub & Xinwei Xiong(cubxxw). All rights reserved.
// Licensed under the MIT License (the "License").

package logrus

import (
	"errors"
	"io"
	"strings"
	"testing"
	"time"

	upstream "github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

type writerEvent struct {
	message string
	level   upstream.Level
}

type writerCapture struct {
	events chan writerEvent
}

func (h writerCapture) Levels() []upstream.Level { return upstream.AllLevels }

func (h writerCapture) Fire(entry *upstream.Entry) error {
	h.events <- writerEvent{message: entry.Message, level: entry.Level}
	return nil
}

func TestHookPreservesFieldsErrorAndCaller(t *testing.T) {
	core, entries := observer.New(zapcore.DebugLevel)
	logger := NewLogger(zap.New(core))
	logger.SetReportCaller(true)
	logger.WithField("contract", "hook").WithError(errors.New("synthetic error")).Warn("synthetic message")
	if entries.Len() != 1 {
		t.Fatalf("expected one actual Zap hook entry, got %d", entries.Len())
	}
	entry := entries.All()[0]
	if entry.Message != "synthetic message" || entry.Level != zapcore.WarnLevel {
		t.Fatalf("hook changed message or level: %v", entry.Entry)
	}
	fields := entry.ContextMap()
	if fields["contract"] != "hook" || fields["error"] != "synthetic error" {
		t.Fatalf("hook lost fields or error: %v", fields)
	}
	if !entry.Caller.Defined {
		t.Fatal("hook lost the actual Logrus caller")
	}
}

func TestWriterContract(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  []string
	}{
		{name: "lines", input: "alpha\nbeta\n", want: []string{"alpha", "beta"}},
		{name: "CRLF and partial EOF", input: "alpha\r\nbeta\r\ntail", want: []string{"alpha", "beta", "tail"}},
		{name: "above scanner limit", input: strings.Repeat("x", 2*65536+7)},
	}
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			core, entries := observer.New(zapcore.DebugLevel)
			logger := NewLogger(zap.New(core))
			logger.SetLevel(upstream.TraceLevel)
			capture := writerCapture{events: make(chan writerEvent, 16)}
			logger.AddHook(capture)
			writer := logger.WithField("contract", "writer").WriterLevel(upstream.ErrorLevel)
			t.Cleanup(func() { _ = writer.Close() })
			done := make(chan error, 1)
			go func() {
				_, err := io.WriteString(writer, test.input)
				closeErr := writer.Close()
				if err == nil {
					err = closeErr
				}
				done <- err
			}()
			select {
			case err := <-done:
				if err != nil {
					t.Fatalf("Writer write/close failed: %v", err)
				}
			case <-time.After(3 * time.Second):
				t.Fatal("Writer write/close timed out")
			}
			var got []string
			deadline := time.After(3 * time.Second)
			for {
				select {
				case event := <-capture.events:
					if event.level != upstream.ErrorLevel {
						t.Fatalf("Writer level changed: %v", event.level)
					}
					got = append(got, event.message)
				case <-deadline:
					t.Fatalf("Writer did not deliver complete content: %d entries", len(got))
				}
				if test.want != nil && len(got) == len(test.want) {
					break
				}
				if test.want == nil && len(strings.Join(got, "")) >= len(test.input) {
					break
				}
			}
			if test.want != nil {
				if strings.Join(got, "|") != strings.Join(test.want, "|") {
					t.Fatalf("Writer changed line/EOF contract: %q", got)
				}
			} else if strings.Join(got, "") != test.input {
				t.Fatal("Writer lost, duplicated or changed long input")
			}
			if entries.Len() != len(got) {
				t.Fatalf("actual Zap hook lost entries: %d vs %d", entries.Len(), len(got))
			}
			for _, entry := range entries.All() {
				if entry.Level != zapcore.ErrorLevel || entry.ContextMap()["contract"] != "writer" {
					t.Fatal("Writer lost level or fields in Zap hook")
				}
			}
		})
	}
}
