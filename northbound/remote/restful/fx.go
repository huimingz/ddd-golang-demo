package restful

import (
	"go.uber.org/fx"

	"demo/northbound/remote/restful/engine"
	"demo/northbound/remote/restful/handler"
	"demo/northbound/remote/restful/middleware"
	"demo/northbound/remote/restful/router"
)

var Module = fx.Module("restful",
	fx.Provide(handler.NewHello),
	fx.Provide(engine.New),
	fx.Provide(router.NewAPIRouter),
	fx.Invoke(middleware.Setup),
	fx.Invoke(router.RegisterHello),
	fx.Invoke(run),
)
