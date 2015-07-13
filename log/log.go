// Copyright (C) 2015 The Android Open Source Project
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package log provides a hierarchical logger interface and implementations of the interface.
package log

// Interface declares the methods for an object that accepts logging messages.
type Interface interface {
	log(severity Severity, msg string, args ...interface{})
	enter(name string) Logger
	fork() Logger
	close()
}

// Severity defines the severity of a logging message.
// The levels match the ones defined in rfc5424 for syslog.
type Severity int

const (
	// Emergency indicates the system is unusable, no further data should be trusted.
	Emergency Severity = 0
	// Alert indicates action must be taken immediately.
	Alert Severity = 1
	// Critical indicates errors severe enough to terminate processing.
	Critical Severity = 2
	// Error indicates non terminal failure conditions that may have an effect on results.
	Error Severity = 3
	// Warning indicates issues that might affect performance or compatibility, but could be ignored.
	Warning Severity = 4
	// Notice indicates normal but significant conditions.
	Notice Severity = 5
	// Info indicates minor informational messages that should generally be ignored.
	Info Severity = 6
	// Debug indicates verbose debug-level messages.
	Debug Severity = 7
)

// Log writes an error message to the logger with the specified severity.
// Arguments are handled in the manner of fmt.Printf.
func Log(l Interface, severity Severity, msg string, args ...interface{}) {
	l.log(severity, msg, args...)
}

// Emergencyf calls l.Log with the Emergency severity.
func Emergencyf(l Interface, msg string, args ...interface{}) {
	l.log(Emergency, msg, args...)
}

// Alertf calls l.Log with the Alert severity.
func Alertf(l Interface, msg string, args ...interface{}) {
	l.log(Alert, msg, args...)
}

// Criticalf calls l.Log with the Critical severity.
func Criticalf(l Interface, msg string, args ...interface{}) {
	l.log(Critical, msg, args...)
}

// Errorf calls l.Log with the Error severity.
func Errorf(l Interface, msg string, args ...interface{}) {
	l.log(Error, msg, args...)
}

// Warningf calls l.Log with the Warning severity.
func Warningf(l Interface, msg string, args ...interface{}) {
	l.log(Warning, msg, args...)
}

// Noticef calls l.Log with the Notice severity.
func Noticef(l Interface, msg string, args ...interface{}) {
	l.log(Notice, msg, args...)
}

// Infof calls l.Log with the Info severity.
func Infof(l Interface, msg string, args ...interface{}) {
	l.log(Info, msg, args...)
}

// Debugf calls l.Log with the Debug severity.
func Debugf(l Interface, msg string, args ...interface{}) {
	l.log(Debug, msg, args...)
}

// E calls l.Log with the Error severity.
func E(l Interface, msg string, args ...interface{}) {
	l.log(Error, msg, args...)
}

// W calls l.Log with the Warning severity.
func W(l Interface, msg string, args ...interface{}) {
	l.log(Warning, msg, args...)
}

// I calls l.Log with the Info severity.
func I(l Interface, msg string, args ...interface{}) {
	l.log(Info, msg, args...)
}

// Enter creates a new logger scoped within the existing logger. This can be used to produce
// hierarchical log messages.
func Enter(l Interface, tag string) Interface {
	return l.enter(tag)
}

// Fork creates a new logger with the same scope as the existing logger, but with a new context
// identifier. It is good practice to fork logs before passing to another goroutine so that
// messages can be associated with their goroutine of execution.
func Fork(l Interface) Interface {
	return l.fork()
}

// Close closes the logger, automatically flushing any remaining messages.
// After calling Close, no other methods can be called on the logger.
func Close(l Interface) {
	l.close()
}
