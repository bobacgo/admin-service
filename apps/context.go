package apps

import "context"

type ctxKey string

const (
	ctxUserIDKey  ctxKey = "uid"
	ctxAccountKey ctxKey = "acc"
)

func WithUserContext(ctx context.Context, userID int64, account string) context.Context {
	ctx = context.WithValue(ctx, ctxUserIDKey, userID)
	ctx = context.WithValue(ctx, ctxAccountKey, account)
	return ctx
}

func UserIDFromContext(ctx context.Context) (int64, bool) {
	v := ctx.Value(ctxUserIDKey)
	if v == nil {
		return 0, false
	}
	id, ok := v.(int64)
	return id, ok
}

func AccountFromContext(ctx context.Context) (string, bool) {
	v := ctx.Value(ctxAccountKey)
	if v == nil {
		return "", false
	}
	s, ok := v.(string)
	return s, ok
}
