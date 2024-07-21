package gin

import ginhttp "ratelimiter/pkg/user/http/ginhttp"

func (s *GinServer) RegisterV1Routes() {
	v1 := s.gin.Group("/api/v1")

	user := v1.Group("/users")
	user.POST("", ginhttp.CreateUser)
	user.GET("", ginhttp.GetUsers)
	user.GET("/:id", ginhttp.GetUser)
	user.PUT("/:id", ginhttp.UpdateUser)
	user.DELETE("/:id", ginhttp.DeleteUser)
}
