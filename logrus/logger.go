// Copyright © 2023 KubeCub & Xinwei Xiong(cubxxw). All rights reserved.
//
// Licensed under the MIT License (the "License");
// you may not use this file except in compliance with the License.

// Package logrus adds a hook to the logrus logger hooks.
package logrus

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

// NewLogger create a logrus logger, add hook to it and return it.
func NewLogger(zapLogger *zap.Logger) *logrus.Logger {
	logger := logrus.New()
	logger.SetOutput(ioutil.Discard)
	logger.AddHook(newHook(zapLogger))

	return logger
}
