package config

// HTTPServer represents the configuration of the http server.
type HTTPServer struct {
	Host      string   `mapstructure:"host" default:"127.0.0.1"`
	Port      int      `mapstructure:"port" default:"8000"`
	Domain    []string `mapstructure:"domain"`
	APIPrefix string   `mapstructure:"api_prefix" default:"/api/v1"`
	Token     string   `mapstructure:"token" default:""`

	// ReverseProxy is a flag to indicate whether this is a reverse proxy, if you
	// use a reverse proxy like nginx, traefik, etc. to proxy the request to
	// this server, you should set this flag to true, otherwise, the real ip of
	// the client may not be able to get.
	ReverseProxy bool `mapstructure:"reverse_proxy" default:"false"`
}
