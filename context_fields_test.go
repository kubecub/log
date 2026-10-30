// Copyright © 2026 KubeCub. All rights reserved.
// Licensed under the MIT License (the "License").

package log

import (
	"context"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

func TestContextFieldsKeepTypedKeysAndClone(t *testing.T) {
	core, entries := observer.New(zapcore.InfoLevel)
	logger := NewLogger(zap.New(core)).(*zapLogger)
	ctx := context.WithValue(context.Background(), KeyRequestID, "synthetic-request")
	ctx = context.WithValue(ctx, KeyUsername, "synthetic-user")
	ctx = context.WithValue(ctx, KeyWatcherName, "synthetic-watcher")
	logger.L(ctx).Info("with typed context")
	logger.Info("base logger")
	got := entries.All()
	if len(got) != 2 {
		t.Fatalf("expected context and base entries, got %d", len(got))
	}
	fields := got[0].ContextMap()
	if len(fields) != 3 || fields["requestID"] != "synthetic-request" || fields["username"] != "synthetic-user" || fields["watcher"] != "synthetic-watcher" {
		t.Fatalf("context field names or values changed: %v", fields)
	}
	if len(got[1].ContextMap()) != 0 {
		t.Fatal("context logging mutated the base logger")
	}
}

func TestPlainStringContextKeysCannotOverrideTypedKeys(t *testing.T) {
	core, entries := observer.New(zapcore.InfoLevel)
	logger := NewLogger(zap.New(core)).(*zapLogger)
	ctx := context.WithValue(context.Background(), "requestID", "synthetic-collision")
	ctx = context.WithValue(ctx, "username", "synthetic-collision")
	ctx = context.WithValue(ctx, "watcher", "synthetic-collision")
	logger.L(ctx).Info("plain string keys")
	got := entries.All()
	if len(got) != 1 || len(got[0].ContextMap()) != 0 {
		t.Fatal("plain string keys crossed the existing typed-key boundary")
	}
}
