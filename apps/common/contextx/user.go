package contextx

import "context"

type userContextKey struct{}

type User struct {
	Account string
	RoleIds []int64
}

func WithUserContext(ctx context.Context, user User) context.Context {
	return context.WithValue(ctx, userContextKey{}, user)
}

func FormUser(ctx context.Context) User {
	user, _ := ctx.Value(userContextKey{}).(User)
	return user
}

func Account(ctx context.Context) string {
	return FormUser(ctx).Account
}

func RoleIds(ctx context.Context) []int64 {
	return FormUser(ctx).RoleIds
}
