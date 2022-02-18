package application

func (server *server) Listen() {
	server.router.GET("", server.auth.Find)
	server.router.Run(":3000")
}
