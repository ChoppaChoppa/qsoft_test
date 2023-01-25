package http

import (
	"github.com/gin-gonic/gin"
	"qsoft_test/internal/qsoft/http/middleware"
)

type IHandler interface {
	GetDays(c *gin.Context)
}

type Server struct {
	*gin.Engine
	middleware *middleware.Middleware
	host       string
	handler    IHandler
}

func New(host string, handler IHandler, middleware *middleware.Middleware) *Server {
	return &Server{
		Engine:     gin.Default(),
		middleware: middleware,
		host:       host,
		handler:    handler,
	}
}

func (s *Server) Run() error {
	initRouters(s)

	if err := s.Engine.Run(s.host); err != nil {
		return err
	}

	return nil
}
