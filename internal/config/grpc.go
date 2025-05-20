package config

type GRPC struct {
	Port int `envconfig:"GRPC_PORT" default:"9000"`
}
