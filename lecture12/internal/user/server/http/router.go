package http

func (s *Server) SetupRoutes() {
	auth := s.App.Group("/api/user/v1")
	auth.GET("/users", s.handler.User.GetAllUsers)
	auth.GET("/:id", s.handler.User.GetUserByID)

}
