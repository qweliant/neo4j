package controllers

import (
	"github.com/tapfunds/tfapi/api/middlewares"
)

func (s *Server) initializeRoutes() {

	v1 := s.Router.Group("/api/v1")
	{
		// API Status
		v1.GET("/status", s.Status)

		// Login Route
		v1.POST("/login", s.Login)

		// Reset password:
		v1.POST("/password/forgot", s.ForgotPassword)
		v1.POST("/password/reset", s.ResetPassword)

		//Users routes
		v1.POST("/users", s.CreateUser)
		// The user of the app have no business getting all the users.
		// v1.GET("/users", s.GetUsers)
		// v1.GET("/users/:id", s.GetUser)
		v1.PUT("/users/:id", middlewares.TokenAuthMiddleware(), s.UpdateUser)
		v1.PUT("/avatar/users/:id", middlewares.TokenAuthMiddleware(), s.UpdateAvatar)
		v1.DELETE("/users/:id", middlewares.TokenAuthMiddleware(), s.DeleteUser)

		//Plaid routes
		v1.POST("/create_link_token", middlewares.TokenAuthMiddleware(), s.createLinkToken)
		v1.POST("/set_access_token", middlewares.TokenAuthMiddleware(), s.getAccessToken)
		v1.POST("/balance", middlewares.TokenAuthMiddleware(), s.balance)
		v1.POST("/auth", middlewares.TokenAuthMiddleware(), s.authorize)
		v1.POST("/accounts", middlewares.TokenAuthMiddleware(), s.accounts)
		v1.POST("/plaid/item", middlewares.TokenAuthMiddleware(), s.item)
		v1.POST("/identity", middlewares.TokenAuthMiddleware(), s.identity)
		v1.POST("/transactions", middlewares.TokenAuthMiddleware(), s.transactions)
		v1.POST("/transfer", middlewares.TokenAuthMiddleware(), s.transfer)

		//Integration Token routes
		v1.GET("/user_integrations/:id", middlewares.TokenAuthMiddleware(),s.GetUserIntegration)
		v1.POST("/new_integration", middlewares.TokenAuthMiddleware(),s.CreatePlaidInfo) // create
		v1.PUT("/integrations/:id", middlewares.TokenAuthMiddleware(),s.UpdateIntegration) // find by id
		v1.DELETE("/integrations/:id", middlewares.TokenAuthMiddleware(),s.DeleteIntegration) // delete by id
	}
}
