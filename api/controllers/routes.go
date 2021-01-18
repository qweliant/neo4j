package controllers

func (s *Server) initializeRoutes() {

	v1 := s.Router.Group("/api/map/v1")
	{
		// API Status
		v1.GET("/status", s.Status)

		//Integration Token routes
		v1.GET("/map_item/:id")
		v1.POST("/map_item") // create
		v1.POST("/query", graphqlHandler())
		v1.GET("/", playgroundHandler())
	}
}
