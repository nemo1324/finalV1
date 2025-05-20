package config

type HTTP struct {
	Port int    `envconfig:"HTTP_PORT" default:"9000"`
	Ip   string `envconfig:"IP" default:"127.0.0.1"`
}
