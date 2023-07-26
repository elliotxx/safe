package safe

import (
	"context"
	"log"
	"runtime/debug"
)

type (
	RecoverHandler func(r interface{})
	DoFunc         func()
	DoCtxFunc      func(ctx context.Context)
)

var DefaultRecoverHandler RecoverHandler = func(r interface{}) {
	log.Printf("Recovered in goroutine: %s, stack: %s\n", r, debug.Stack())
}

// Go starts a recoverable goroutine.
//
// Example:
//
//	safe.Go(
//	  func() { ... },
//	)
func Go(do DoFunc) {
	GoR(do, DefaultRecoverHandler)
}

// GoR starts a recoverable goroutine using given recover handler.
//
// Example:
//
//	safe.GoR(
//	  func() { ... },
//	  customRecoverHandler
//	)
func GoR(do DoFunc, handlers ...RecoverHandler) {
	go func() {
		defer HandleCrash(handlers...)
		do()
	}()
}

// DefaultHandleCrash simply catches a crash with the default recover handler.
//
// Example:
//
//	go func() {
//	  defer DefaultHandleCrash()
//	  ...
//	}()
func DefaultHandleCrash() {
	HandleCrash(DefaultRecoverHandler)
}

// HandleCrash catches a crash with the custom recover handlers.
//
// Example:
//
//	go func() {
//	  defer HandleCrash(customRecoverHandler)
//	  ...
//	}()
func HandleCrash(handlers ...RecoverHandler) {
	if r := recover(); r != nil {
		for _, handler := range handlers {
			handler(r)
		}
	}
}
