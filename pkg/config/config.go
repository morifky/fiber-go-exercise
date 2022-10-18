package config

type Config struct {
	DBHost     string `env:"DB_HOST" envDefault:"2"`
	DBPort     string `env:"DB_PORT" envDefault:"2"`
	DBUsername string `env:"DB_USERNAME" envDefault:"1"`
	DBPassword string `env:"DB_PASSWORD" envDefault:"1"`
	DBName     string `env:"DB_NAME" envDefault:"1"`
	HTTPPort   string `env:"PORT" envDefault:"8080"`
}

func InitConfig() *Config {
	return &Config{}
}
