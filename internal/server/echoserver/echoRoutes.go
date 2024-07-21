package echo

import echohttp "ratelimiter/pkg/user/http/echohttp"

func (s *EchoServer) RegisterV1Routes() {
	v1 := s.Echo.Group("/api/v1")

	user := v1.Group("/users")
	user.POST("", echohttp.CreateUser)
	user.GET("", echohttp.GetUsers)
	user.GET("/:id", echohttp.GetUser)
	user.PUT("/:id", echohttp.UpdateUser)
	user.DELETE("/:id", echohttp.DeleteUser)
}
