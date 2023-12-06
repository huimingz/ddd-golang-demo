package configloader

import (
	"context"

	"demo/config"
	"demo/extension/cfg"
	"demo/extension/contextz"
)

func FromYaml(ctx context.Context) *config.Schema {
	var conf config.Schema
	directory := contextz.ConfigDirectory(ctx, "./etc/")
	fileName := contextz.ConfigFilename(ctx, "config.yaml")

	cfg.MustLoad(&conf, &cfg.Options{ConfigPath: directory, ConfigName: fileName})
	return &conf
}
