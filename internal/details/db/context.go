package db

import (
	"context"
)

type ConnContext struct {
	Resolver func(ctx context.Context) Executor
}

type TxContext struct {
	Tx Tx
}

type ctxConnContextKey struct{}
type ctxTxContextKey struct{}

func WithConnContext(ctx context.Context, v ConnContext) context.Context {
	ctx = context.WithValue(ctx, ctxConnContextKey{}, v)

	return ctx
}

func GetConnContext(ctx context.Context) (ConnContext, bool) {
	v, ok := ctx.Value(ctxConnContextKey{}).(ConnContext)
	return v, ok
}

func WithTxContext(ctx context.Context, v TxContext) context.Context {
	ctx = context.WithValue(ctx, ctxTxContextKey{}, v)

	return ctx
}

func GetTxContext(ctx context.Context) (TxContext, bool) {
	v, ok := ctx.Value(ctxTxContextKey{}).(TxContext)
	return v, ok
}
