package safe

import (
	"testing"
)

func TestGoroutineRecover(t *testing.T) {
	// If recover fails the test will panic
	Go(func() {
		panic("BOOM")
	})
}
