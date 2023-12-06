package remote

import (
	"go.uber.org/fx"

	"demo/northbound/remote/restful"
	"demo/southbound/adapter/configloader"
)

var Module = fx.Module("remote",
	fx.Provide(configloader.FromYaml),
	restful.Module,
)
