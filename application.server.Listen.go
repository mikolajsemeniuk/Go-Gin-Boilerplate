package application

func (s *server) Listen() {
	s.router.GET("", s.auth.Find)
	s.router.Run(":3000")
}
