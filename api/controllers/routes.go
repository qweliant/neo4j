package controllers

func (s *Server) initializeRoutes() {

	v1 := s.Router.Group("/api/map/v1")
	{
		// API Status
		v1.GET("/status", s.Status)

		//Object handling routes
		v1.GET("/map_item/:id")
		v1.POST("/map_item", s.CreateUserItem) // create
	}
}