package contextz

import (
	"context"
)

var moduleKey = &contextKey{name: "module"}

func Clone(ctx context.Context) context.Context {
	return ctx
}

func AsyncClone(ctx context.Context) context.Context {
	rv := context.Background()
	moduleName := ModuleName(ctx)
	if moduleName != "" {
		rv = WithModuleName(rv, moduleName)
	}

	// span := tracing.SpanFromContext(ctx)
	// if !span.Empty() {
	// 	rv = tracing.ContextWithSpan(rv, span)
	// }
	return rv
}

func WithModuleName(ctx context.Context, moduleName string) context.Context {
	return context.WithValue(ctx, moduleKey, moduleName)
}

func ModuleName(ctx context.Context) string {
	moduleName, _ := ctx.Value(moduleKey).(string)
	return moduleName
}
