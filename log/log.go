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
	// Log writes an error message to the logger with the specified severity.
	// Arguments are handled in the manner of fmt.Printf.
	Log(severity Severity, msg string, args ...interface{})
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

// Emergencyf calls l.Log with the Emergency severity.
func Emergencyf(l Interface, msg string, args ...interface{}) {
	l.Log(Emergency, msg, args...)
}

// Alertf calls l.Log with the Alert severity.
func Alertf(l Interface, msg string, args ...interface{}) {
	l.Log(Alert, msg, args...)
}

// Criticalf calls l.Log with the Critical severity.
func Criticalf(l Interface, msg string, args ...interface{}) {
	l.Log(Critical, msg, args...)
}

// Errorf calls l.Log with the Error severity.
func Errorf(l Interface, msg string, args ...interface{}) {
	l.Log(Error, msg, args...)
}

// Warningf calls l.Log with the Warning severity.
func Warningf(l Interface, msg string, args ...interface{}) {
	l.Log(Warning, msg, args...)
}

// Noticef calls l.Log with the Notice severity.
func Noticef(l Interface, msg string, args ...interface{}) {
	l.Log(Notice, msg, args...)
}

// Infof calls l.Log with the Info severity.
func Infof(l Interface, msg string, args ...interface{}) {
	l.Log(Info, msg, args...)
}

// Debugf calls l.Log with the Debug severity.
func Debugf(l Interface, msg string, args ...interface{}) {
	l.Log(Debug, msg, args...)
}

// E calls l.Log with the Error severity.
func E(l Interface, msg string, args ...interface{}) {
	l.Log(Error, msg, args...)
}

// W calls l.Log with the Warning severity.
func W(l Interface, msg string, args ...interface{}) {
	l.Log(Warning, msg, args...)
}

// I calls l.Log with the Info severity.
func I(l Interface, msg string, args ...interface{}) {
	l.Log(Info, msg, args...)
}
