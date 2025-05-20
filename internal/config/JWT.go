package config

type JWT struct {
	Secret string `envconfig:"SECRET"`
}
