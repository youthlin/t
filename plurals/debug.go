package plurals

import "context"

// ctxKey custom type using by context.WithValue as a key
type ctxKey string

// ctxKeyDebug a debug key instace
var ctxKeyDebug = ctxKey("debug")

// DebugCtx return a ctx that with a debug key, which may print how the exp calculates.
func DebugCtx(ctx context.Context) context.Context {
	return context.WithValue(ctx, ctxKeyDebug, true)
}
