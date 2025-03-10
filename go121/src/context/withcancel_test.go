package context_test

import (
	"context"
	"errors"
	"testing"
)

// WithoutCancel function returns a copy of a context that is not canceled when the original context is canceled.
func TestWithCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	t.Log(ctx.Err(), context.Cause(ctx))

	cancel()
	t.Log(ctx.Err(), context.Cause(ctx))
	t.Log(errors.Is(ctx.Err(), context.Canceled), errors.Is(context.Cause(ctx), context.Canceled))

	r, _ := context.WithCancelCause(ctx)
	t.Log(r.Err(), context.Cause(r)) // already canceled
}

func TestWithoutCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	t.Log(ctx.Err(), context.Cause(ctx))

	cancel()
	t.Log(ctx.Err(), context.Cause(ctx))
	t.Log(errors.Is(ctx.Err(), context.Canceled), errors.Is(context.Cause(ctx), context.Canceled))

	r := context.WithoutCancel(ctx)
	t.Log(r.Err(), context.Cause(r)) // not canceled yet
}
