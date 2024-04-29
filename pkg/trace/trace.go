// Package trace provides utilities for error tracing.
package trace

import (
	"fmt"
	"runtime"
	"strings"
)

func funcName(skip int) string {
	pc, _, _, ok := runtime.Caller(skip)
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
	return fmt.Errorf(funcName(option.Skip)+": %w", err)
}

// WithSkip sets the number of stack frames to skip when identifying the caller.
func WithSkip(skip int) WrapOption {
	return func(wo *WrapOpt) {
		wo.Skip = skip + defaultSkip
	}
}
