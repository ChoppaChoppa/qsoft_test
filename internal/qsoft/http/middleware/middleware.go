package middleware

import "github.com/gin-gonic/gin"

type Middleware struct {
	XPong string
}

func New(xPong string) *Middleware {
	return &Middleware{
		XPong: xPong,
	}
}

func (m *Middleware) CheckPing(c *gin.Context) {
	ping := c.GetHeader("X-PING")
	if ping != "ping" {
		c.Next()
		return
	}

	c.Header("X-PONG", m.XPong)
	c.Next()
}
