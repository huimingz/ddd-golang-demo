package main

import (
	"context"

	"go.uber.org/fx"

	"demo/northbound/remote"
)

func main() {
	app := fx.New(
		fx.Provide(context.TODO),
		remote.Module,
	)
	app.Run()
}
