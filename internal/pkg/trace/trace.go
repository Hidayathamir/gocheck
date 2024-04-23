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

// Wrap -.
func Wrap(err error, options ...WrapOption) error {
	opt := &WrapOpt{}
	for _, o := range options {
		o(opt)
	}
	if opt.Skip == 0 {
		opt.Skip = 2
	}
	return fmt.Errorf(funcName(opt.Skip)+": %w", err)
}

// WithSkip -.
func WithSkip(skip int) WrapOption {
	return func(wo *WrapOpt) {
		wo.Skip = skip + 2 //nolint:gomnd
	}
}
