// Copyright © 2023 KubeCub. All rights reserved.
//
// Licensed under the MIT License (the "License");
// you may not use this file except in compliance with the License.

package main

import (
	"log"
	"os"

	"github.com/kubecub/log/cuslog"
)

func main() {
	cuslog.Info("std log")
	cuslog.SetOptions(cuslog.WithLevel(cuslog.InfoLevel))
	cuslog.Debug("change std log to debug level")

	// 输出到文件
	fd, err := os.OpenFile("test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln("create file test.log failed")
	}
	defer fd.Close()

	l := cuslog.New(
		cuslog.WithLevel(cuslog.DebugLevel),
		cuslog.WithOutput(fd),
		cuslog.WithFormatter(&cuslog.JsonFormatter{IgnoreBasicFields: false}),
		//cuslog.WithFormatter(&cuslog.TextFormatter{IgnoreBasicFields: false}),
	)
	l.Debugf("custom log with json formatter: %s", "test")
	l.Error("custom log with json formatter")
	l.Info("custom log with json formatter") // Only info will be logged
}
