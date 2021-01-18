package main

import (
	"fmt"

	"github.com/mindstand/gogm"
	"github.com/qweliant/neo4j/api/models"
)

func main() {
	config := gogm.Config{
		IndexStrategy: gogm.VALIDATE_INDEX, //other options are ASSERT_INDEX and IGNORE_INDEX
		PoolSize:      50,
		Port:          7687,
		IsCluster:     false, //tells it whether or not to use `bolt+routing`
		Host:          "0.0.0.0",
		Password:      "DB_DRIVER",
		Username:      "DB_USER",
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
	sess, err := gogm.NewSession(false)
	if err != nil {
		panic(err)
	}

	//close the session
	defer sess.Close()
 
	fmt.Println("Done")
}
