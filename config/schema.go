package config

type Schema struct {
	Name string     `mapstructure:"name" default:"demo"`
	HTTP HTTPServer `mapstructure:"http"`
	CORS CORSConfig `mapstructure:"cors"`
}
