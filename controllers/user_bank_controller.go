package controller
//// create user object in DB


//// client gets account info vis identity plaid endpoint from plaid service after linking
//// client will then send info to backend
//// user model gets created on server
//// client will query db for user with matching ID and return:
//// 	item model
import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"	
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/qweliant/neo4j/models"

)

func 