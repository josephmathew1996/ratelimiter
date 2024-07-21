package gin

import (
	"github.com/gin-gonic/gin"
)

type GinServer struct {
	gin *gin.Engine
}

func NewGinServer() *GinServer {
	g := gin.Default()

	return &GinServer{
		gin: g,
	}
}

func (s *GinServer) Start() {
	s.gin.Run(":8081")
}
