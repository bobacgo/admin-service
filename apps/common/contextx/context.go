package contextx

import "context"

type ipContextKey struct{}

func WithIPContext(ctx context.Context, ip string) context.Context {
	return context.WithValue(ctx, ipContextKey{}, ip)
}

func IP(ctx context.Context) string {
	ip, _ := ctx.Value(ipContextKey{}).(string)
	return ip
}
