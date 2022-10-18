package safe

import (
	"context"
	"sync"
)

// Pool is a pool of go routines.
type Pool struct {
	waitGroup sync.WaitGroup
	ctx       context.Context
	cancel    context.CancelFunc
}

// NewPool creates a Pool.
func NewPool(parentCtx context.Context) *Pool {
	ctx, cancel := context.WithCancel(parentCtx)
	return &Pool{
		ctx:    ctx,
		cancel: cancel,
	}
}

// GoCtx starts a recoverable goroutine with a context.
func (p *Pool) GoCtx(do DoCtxFunc) {
	p.waitGroup.Add(1)
	Go(func() {
		defer p.waitGroup.Done()
		do(p.ctx)
	})
}

// Stop stops all started routines, waiting for their termination.
func (p *Pool) Stop() {
	p.cancel()
	p.waitGroup.Wait()
}

// Wait waits all started routines.
func (p *Pool) Wait() {
	p.waitGroup.Wait()
}
