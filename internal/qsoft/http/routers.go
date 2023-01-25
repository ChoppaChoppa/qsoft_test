package http

import "github.com/gin-gonic/gin"

func initRouters(s *Server) {
	s.Engine.Use(gin.Logger())
	s.Engine.Use(s.middleware.CheckPing)

	initDaysRouter(s)
}

func initDaysRouter(s *Server) {
	s.Engine.POST("/when/:year", s.handler.GetDays)
}
