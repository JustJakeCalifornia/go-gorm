package config

type Config struct {
	WebPort string
}

func NewConfig() *Config {
	return &Config{
		WebPort: "8080",
	}
}
