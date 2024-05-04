// Package trace provides utilities for error tracing.
package trace

import (
	"fmt"
	"runtime"
	"strings"
)

type optFuncName struct {
	Skip int
}

type optionFuncName func(*optFuncName)

// FuncName return caller function name.
func FuncName(options ...optionFuncName) string {
	option := &optFuncName{
		Skip: 1,
	}
	for _, opt := range options {
		opt(option)
	}

	pc, _, _, ok := runtime.Caller(option.Skip)
	if !ok {
		return "?"
	}

	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return "?"
	}

	funcNameWithModule := fn.Name()
	funcNameWithModuleSplit := strings.Split(funcNameWithModule, "/")
	funcName := funcNameWithModuleSplit[len(funcNameWithModuleSplit)-1]

	return funcName
}

func withSkipFuncName(skip int) optionFuncName {
	return func(o *optFuncName) {
		o.Skip = skip
	}
}

// WrapOpt represents options for the Wrap function.
type WrapOpt struct {
	Skip int
}

// WrapOption represents an option for the Wrap function.
type WrapOption func(*WrapOpt)

const defaultSkip = 2

// Wrap wraps the given error with the name of the calling function.
func Wrap(err error, options ...WrapOption) error {
	option := &WrapOpt{
		Skip: defaultSkip,
	}
	for _, opt := range options {
		opt(option)
	}
	return fmt.Errorf(FuncName(withSkipFuncName(option.Skip))+": %w", err)
}

// WithSkip sets the number of stack frames to skip when identifying the caller.
func WithSkip(skip int) WrapOption {
	return func(wo *WrapOpt) {
		wo.Skip = skip + defaultSkip
	}
}
