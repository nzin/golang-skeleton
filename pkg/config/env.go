package config

// Config is the whole configuration of the app
var Config = struct {
	// Host - golang-skeleton server host
	Host string `env:"HOST" envDefault:"localhost"`
	// Port - golang-skeleton server port
	Port int `env:"PORT" envDefault:"18000"`

	// LogrusLevel sets the logrus logging level
	LogrusLevel string `env:"GOLANG_SKELETON_LOGRUS_LEVEL" envDefault:"info"`
	// LogrusFormat sets the logrus logging formatter
	// Possible values: text, json
	LogrusFormat string `env:"GOLANG_SKELETON_LOGRUS_FORMAT" envDefault:"json"`

	// MiddlewareVerboseLoggerEnabled - to enable the negroni-logrus logger for all the endpoints useful for debugging
	MiddlewareVerboseLoggerEnabled bool `env:"GOLANG_SKELETON_MIDDLEWARE_VERBOSE_LOGGER_ENABLED" envDefault:"true"`
	// MiddlewareVerboseLoggerExcludeURLs - to exclude urls from the verbose logger via comma separated list
	MiddlewareVerboseLoggerExcludeURLs []string `env:"GOLANG_SKELETON_MIDDLEWARE_VERBOSE_LOGGER_EXCLUDE_URLS" envDefault:"" envSeparator:","`
	// MiddlewareGzipEnabled - to enable gzip middleware
	MiddlewareGzipEnabled bool `env:"GOLANG_SKELETON_MIDDLEWARE_GZIP_ENABLED" envDefault:"true"`
}{}
