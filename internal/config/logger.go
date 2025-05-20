package config

import (
	"final/internal/utils/observability/log"
)

type Logger struct {
	Level log.Level `envconfig:"LOG_LEVEL" default:"0"`
}
