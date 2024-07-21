package main

import (
	"log"
	echo "ratelimiter/internal/server/echoserver"
)

func main() {
	//load
	echoServer, err := echo.NewEchoServer()
	if err != nil {
		log.Fatalf("Failed to initialize server: %v", err)
	}
	echoServer.RegisterV1Routes()
	echoServer.Start()

	// ginServer := gin.NewGinServer()
	// ginServer.RegisterV1Routes()
	// ginServer.Start()
}
