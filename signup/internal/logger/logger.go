package logger

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

type Logger interface {
	LogConnectionInfo(r *http.Request)
	LogError(err error)
}

type ZeroLogger struct {
}

func NewZeroLogger() *ZeroLogger {
	return &ZeroLogger{}
}

func (*ZeroLogger) LogConnectionInfo(r *http.Request) {
	log.Info().
		Str("method", r.Method).
		Str("path", r.URL.Path).
		Msg("received request: ")
}

func (*ZeroLogger) LogError(err error) {
	log.Error().
		Msg(err.Error())
}
