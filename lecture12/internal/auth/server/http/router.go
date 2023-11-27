package http

func (s *Server) SetupRoutes() {
	auth := s.App.Group("/api/auth/v1")
	auth.POST("/sign-in", s.handler.UserToken.Login)
}
