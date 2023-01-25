package main

import (
	"github.com/rs/zerolog"
	"os"
	"os/signal"
	"qsoft_test/internal/qsoft"
	"qsoft_test/internal/qsoft/config"
	"qsoft_test/internal/qsoft/http"
	"qsoft_test/internal/qsoft/http/handler"
	"qsoft_test/internal/qsoft/http/middleware"
	"syscall"
	"time"
)

func main() {
	out := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.StampMilli,
	}

	logger := zerolog.New(out).With().Caller().Logger().With().Timestamp().Logger()

	cfg, err := config.Parse()
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to parse configs")
	}

	svc := qsoft.New(logger)

	h := handler.New(logger, svc)
	m := &middleware.Middleware{XPong: "pong"}

	server := http.New(cfg.Server.Host, h, m)
	go func() {
		if err = server.Run(); err != nil {
			logger.Fatal().Err(err).Msg("failed to parse configs")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	//logger.Info().Msg("http server shutdown")
	//
	//if err = server.Engine.Shu; err != nil {
	//	logger.Fatal().Err(err).Msg("server shutdown error")
	//}
}
