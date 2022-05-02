package configuration

type Config struct {
	Log
}

type Log struct {
	Level string `env:"LOG_LEVEL" envDefault:"DEBUG"`
}
