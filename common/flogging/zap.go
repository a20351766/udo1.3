/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package flogging

import (
	"fmt"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zapgrpc"
)

// NewZapLogger creates a zap logger around a new zap.Core. The core will use
// the provided encoder and sinks and a level enabler that is associated with
// the provided module name. The logger that is returned will be named the same
// as the module.
func NewZapLogger(core zapcore.Core, options ...zap.Option) *zap.Logger {
	return zap.New(
		core,
		append([]zap.Option{
			zap.AddCaller(),
			zap.AddStacktrace(zapcore.ErrorLevel),
		}, options...)...,
	)
}

// NewGRPCLogger creates a grpc.Logger that delegates to a zap.Logger.
func NewGRPCLogger(l *zap.Logger) *zapgrpc.Logger {
	l = l.WithOptions(
		zap.AddCaller(),
		zap.AddCallerSkip(3),
	)
	return zapgrpc.NewLogger(l, zapgrpc.WithDebug())
}

// NewUdoLogger creates a logger that delegates to the zap.SugaredLogger.
func NewUdoLogger(l *zap.Logger, options ...zap.Option) *UdoLogger {
	return &UdoLogger{
		s: l.WithOptions(append(options, zap.AddCallerSkip(1))...).Sugar(),
	}
}

// A UdoLogger is an adapter around a zap.SugaredLogger that provides
// structured logging capabilities while preserving much of the legacy logging
// behavior.
//
// The most significant difference between the UdoLogger and the
// zap.SugaredLogger is that methods without a formatting suffix (f or w) build
// the log entry message with fmt.Sprintln instead of fmt.Sprint. Without this
// change, arguments are not separated by spaces.
type UdoLogger struct{ s *zap.SugaredLogger }

func (f *UdoLogger) DPanic(args ...interface{})                    { f.s.DPanicf(formatArgs(args)) }
func (f *UdoLogger) DPanicf(template string, args ...interface{})  { f.s.DPanicf(template, args...) }
func (f *UdoLogger) DPanicw(msg string, kvPairs ...interface{})    { f.s.DPanicw(msg, kvPairs...) }
func (f *UdoLogger) Debug(args ...interface{})                     { f.s.Debugf(formatArgs(args)) }
func (f *UdoLogger) Debugf(template string, args ...interface{})   { f.s.Debugf(template, args...) }
func (f *UdoLogger) Debugw(msg string, kvPairs ...interface{})     { f.s.Debugw(msg, kvPairs...) }
func (f *UdoLogger) Error(args ...interface{})                     { f.s.Errorf(formatArgs(args)) }
func (f *UdoLogger) Errorf(template string, args ...interface{})   { f.s.Errorf(template, args...) }
func (f *UdoLogger) Errorw(msg string, kvPairs ...interface{})     { f.s.Errorw(msg, kvPairs...) }
func (f *UdoLogger) Fatal(args ...interface{})                     { f.s.Fatalf(formatArgs(args)) }
func (f *UdoLogger) Fatalf(template string, args ...interface{})   { f.s.Fatalf(template, args...) }
func (f *UdoLogger) Fatalw(msg string, kvPairs ...interface{})     { f.s.Fatalw(msg, kvPairs...) }
func (f *UdoLogger) Info(args ...interface{})                      { f.s.Infof(formatArgs(args)) }
func (f *UdoLogger) Infof(template string, args ...interface{})    { f.s.Infof(template, args...) }
func (f *UdoLogger) Infow(msg string, kvPairs ...interface{})      { f.s.Infow(msg, kvPairs...) }
func (f *UdoLogger) Panic(args ...interface{})                     { f.s.Panicf(formatArgs(args)) }
func (f *UdoLogger) Panicf(template string, args ...interface{})   { f.s.Panicf(template, args...) }
func (f *UdoLogger) Panicw(msg string, kvPairs ...interface{})     { f.s.Panicw(msg, kvPairs...) }
func (f *UdoLogger) Warn(args ...interface{})                      { f.s.Warnf(formatArgs(args)) }
func (f *UdoLogger) Warnf(template string, args ...interface{})    { f.s.Warnf(template, args...) }
func (f *UdoLogger) Warnw(msg string, kvPairs ...interface{})      { f.s.Warnw(msg, kvPairs...) }
func (f *UdoLogger) Warning(args ...interface{})                   { f.s.Warnf(formatArgs(args)) }
func (f *UdoLogger) Warningf(template string, args ...interface{}) { f.s.Warnf(template, args...) }

// for backwards compatibility
func (f *UdoLogger) Critical(args ...interface{})                   { f.s.Errorf(formatArgs(args)) }
func (f *UdoLogger) Criticalf(template string, args ...interface{}) { f.s.Errorf(template, args...) }
func (f *UdoLogger) Notice(args ...interface{})                     { f.s.Infof(formatArgs(args)) }
func (f *UdoLogger) Noticef(template string, args ...interface{})   { f.s.Infof(template, args...) }

func (f *UdoLogger) Named(name string) *UdoLogger { return &UdoLogger{s: f.s.Named(name)} }
func (f *UdoLogger) Sync() error                     { return f.s.Sync() }

func (f *UdoLogger) IsEnabledFor(level zapcore.Level) bool {
	return f.s.Desugar().Core().Enabled(level)
}

func (f *UdoLogger) With(args ...interface{}) *UdoLogger {
	return &UdoLogger{s: f.s.With(args...)}
}

func (f *UdoLogger) WithOptions(opts ...zap.Option) *UdoLogger {
	l := f.s.Desugar().WithOptions(opts...)
	return &UdoLogger{s: l.Sugar()}
}

func formatArgs(args []interface{}) string { return strings.TrimSuffix(fmt.Sprintln(args...), "\n") }
