package controllers

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mindstand/gogm"

	"github.com/qweliant/neo4j/api/middlewares"
	"github.com/qweliant/neo4j/api/models"
)

type Server struct {
	DB     *gogm.Session
	Router *gin.Engine
}

func (server *Server) Database(DbUser string, DbPassword string, DbHost string, DbPort int) {
	config := gogm.Config{
		IndexStrategy: gogm.VALIDATE_INDEX, //other options are ASSERT_INDEX and IGNORE_INDEX
		PoolSize:      50,
		Port:          DbPort,
		IsCluster:     false, //tells it whether or not to use `bolt+routing`
		Host:          DbHost,
		Password:      DbPassword,
		Username:      DbUser,
	}

	err := gogm.Init(
		&config,
		&models.User{},
		&models.Item{},
		&models.Institution{},
		&models.Account{},
		&models.Balance{},
		&models.Owner{},
		&models.Name{},
		&models.PhoneNumber{},
		&models.Email{},
		&models.Address{},
		&models.Location{},
		&models.Transaction{},
	)
	if err != nil {
		panic(err)
	}

	//param is readonly, we're going to make stuff so we're going to do read write
	server.DB, err = gogm.NewSession(false)
	if err != nil {
		panic(err)
	}

	server.Router = gin.Default()
	server.Router.Use(middlewares.CORSMiddleware())

	server.initializeRoutes()
	
	//close the session
	defer server.DB.Close()
}

func (server *Server) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

func (server *Server) Status(c *gin.Context) {

	_, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"error":  "The world has falllen and we are to slumber...",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
	})
}
