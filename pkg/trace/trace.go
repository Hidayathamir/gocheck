// Package trace -.
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

// WrapOpt -.
type WrapOpt struct {
	Skip int
}

// WrapOption -.
type WrapOption func(*WrapOpt)

const defaultSkip = 2

// Wrap -.
func Wrap(err error, options ...WrapOption) error {
	option := &WrapOpt{
		Skip: defaultSkip,
	}
	for _, opt := range options {
		opt(option)
	}
	return fmt.Errorf(funcName(option.Skip)+": %w", err)
}

// WithSkip -.
func WithSkip(skip int) WrapOption {
	return func(wo *WrapOpt) {
		wo.Skip = skip + defaultSkip
	}
}
