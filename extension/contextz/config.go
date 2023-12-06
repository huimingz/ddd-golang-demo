package contextz

import "context"

var (
	// d := ctx.Value("config.directory")
	// n := ctx.Value("config.filename")
	configDirectoryContextKey = &contextKey{name: "config.directory"}
	configFilenameContextKey  = &contextKey{name: "config.filename"}
)

func WithConfigDirectory(ctx context.Context, dir string) context.Context {
	return context.WithValue(ctx, configDirectoryContextKey, dir)
}

func ConfigDirectory(ctx context.Context, defaultValue string) string {
	d := ctx.Value(configDirectoryContextKey)
	if v, ok := d.(string); ok {
		return v
	}
	return defaultValue
}

func WithConfigFilename(ctx context.Context, name string) context.Context {
	return context.WithValue(ctx, configFilenameContextKey, name)
}

func ConfigFilename(ctx context.Context, defaultValue string) string {
	n := ctx.Value(configFilenameContextKey)
	if v, ok := n.(string); ok {
		return v
	}
	return ""
}
