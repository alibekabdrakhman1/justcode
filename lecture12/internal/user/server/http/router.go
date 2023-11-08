package http

func (s *Server) SetupRoutes() {
	auth := s.App.Group("/api/user/v1", s.m.ValidateAuth)
	auth.GET("/users", s.handler.User.GetAllUsers)
	auth.GET("/:id", s.handler.User.GetUserByID)

}
