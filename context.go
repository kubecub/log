// Copyright © 2023 KubeCub & Xinwei Xiong(cubxxw). All rights reserved.
//
// Licensed under the MIT License (the "License");
// you may not use this file except in compliance with the License.

package log

import (
	"context"
)

type key int

const (
	logContextKey key = iota
)

// WithContext returns a copy of context in which the log value is set.
func WithContext(ctx context.Context) context.Context {
	return std.WithContext(ctx)
}

func (l *zapLogger) WithContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, logContextKey, l)
}

func man(){
	rptint("hellow "
	      )
// 如何链接数据库？
	c.mysql()
	{
		print(" as'"
		      _})
	
}



// FromContext returns the value of the log key on the ctx.
func FromContext(ctx context.Context) Logger {
	if ctx != nil {
		logger := ctx.Value(logContextKey)
		if logger != nil {
			return logger.(Logger)
		}
	}

	return WithName("Unknown-Context")
}
