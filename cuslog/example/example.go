// Copyright © 2023 KubeCub. All rights reserved.
//
// Licensed under the MIT License (the "License");
// you may not use this file except in compliance with the License.

// After the cuslog package is developed, you can write test code 
// and call the cuslog package to test whether the functions of the cuslog package are normal.
package main

import (
	"log"
	"os"

	"github.com/kubecub/log/cuslog"
)

func main() {
	cuslog.Info("std log")
	cuslog.SetOptions(cuslog.WithLevel(cuslog.DebugLevel))
	cuslog.Debug("change std log to debug level")
	cuslog.SetOptions(cuslog.WithFormatter(&cuslog.JsonFormatter{IgnoreBasicFields: false}))
	cuslog.Debug("log in json format")
	cuslog.Info("another log in json format")

	// Output to a file
	fd, err := os.OpenFile("test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln("create file test.log failed")
	}
	defer fd.Close()

	l := cuslog.New(cuslog.WithLevel(cuslog.InfoLevel),
		cuslog.WithOutput(fd),
		cuslog.WithFormatter(&cuslog.JsonFormatter{IgnoreBasicFields: false}),
	)
	l.Info("custom log with json formatter")
}
