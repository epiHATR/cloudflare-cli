package config

type Config struct {
	Auth auth
}

type auth struct {
	Token string `mapstructure:"token"`
	Email string `mapstructure:"email"`
	Key   string `mapstructure:"key"`
}
