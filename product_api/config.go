package product_api

type Config struct {
	AppName  string `env:"APP_NAME" env-default:"PRODUCT"`
	AppEnv   string `env:"APP_ENV" env-default:"Development"`
	Port     string `env:"APP_PORT" env-default:"9090"`
	Host     string `env:"HOST" env-default:"localhost"`
	LogLevel string `env:"LOG_LEVEL" env-default:"ERROR"`
}

var (
	cfg Config
)
