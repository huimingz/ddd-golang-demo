package config

type CORSConfig struct {
	Enable          bool `mapstructure:"enable" default:"false"`
	AllowAllOrigins bool `mapstructure:"allow_all_origins" default:"true"`

	// AllowOrigins is a list of origins a cross-domain request can be executed from.
	// If the special "*" value is present in the list, all origins will be allowed.
	// Default value is []
	AllowOrigins []string `mapstructure:"allow_origins"`

	// AllowMethods is a list of methods the client is allowed to use with
	// cross-domain requests. Default value is simple methods (GET and POST)
	AllowMethods []string `mapstructure:"allow_methods"`

	// AllowHeaders is list of non simple headers the client is allowed to use with
	// cross-domain requests.
	AllowHeaders []string `mapstructure:"allow_headers"`

	// AllowCredentials indicates whether the request can include user credentials like
	// cookies, HTTP authentication or client side SSL certificates.
	AllowCredentials bool `mapstructure:"allow_credentials" default:"false"`

	// ExposedHeaders indicates which headers are safe to expose to the API of a CORS
	// API specification
	ExposeHeaders []string `mapstructure:"expose_headers"`

	// MaxAge indicates how long (in seconds) the results of a preflight request
	// can be cached
	MaxAge uint64 `mapstructure:"max_age_seconds" default:"43200"`

	// Allows to add origins like http://some-domain/*, https://api.* or http://some.*.subdomain.com
	AllowWildcard bool `mapstructure:"allow_wildcard" default:"false"`
}
