package config

type Postgres struct {
	Host     string `envconfig:"POSTGRES_HOST" required:"true"`
	Port     string `envconfig:"POSTGRES_PORT" required:"true"`
	User     string `envconfig:"POSTGRES_USER" required:"true"`
	DbName   string `envconfig:"POSTGRES_DB" required:"true"`
	Password string `envconfig:"POSTGRES_PASSWORD" required:"true"`
	SSLMode  string `envconfig:"POSTGRES_SSLMODE" default:"disable"`
}
