package main

import (
	"fmt"
	"time"

	"github.com/mindstand/gogm"
)

type tdString string
type tdInt int

type Institution struct {
	// provides required node fields
	gogm.BaseNode
	IntstitutionID string   `gogm:"name=inst_id"`
	Name           string   `gogm:"name=name"`
	Products       []string `gogm:"name=prods"`
	PrimaryColor   string   `gogm:"name=prim_color"`
	Logo           string   `gogm:"name=logo"`
	Accounts       []*Account `gogm:"direction=outgoing;relationship=account"`
}

type Account struct {
	gogm.BaseNode
	AccountID          string       `gogm:"name=accnt_id"`
	Name               string       `gogm:"name=name"`
	OfficialName       string       `gogm:"name=offic_name"`
	Type               string       `gogm:"name=type"`
	Subtype            string       `gogm:"name=subtype"`
	VerificationStatus bool         
	Owner              *Owner       `gogm:"direction=outgoing;relationship=owner"`
	Balance            *Balance     `gogm:"direction=outgoing;relationship=balance"`
	Institution        *Institution `gogm:"direction=incoming;relationship=account"`
}

type Balance struct {
	gogm.BaseNode
	Available float64   `gogm:"name=avail"`
	Current   float64	`gogm:"name=current"`
	Limit     float64   `gogm:"name=limit"`
	Currency  string	`gogm:"name=currency"`
	Accounts  *Account  `gogm:"direction=incoming;relationship=account"`
}

type Owner struct {
	gogm.BaseNode
	Names        []*Name        `gogm:"direction=outgoing;relationship=name"`
	PhoneNumbers []*PhoneNumber `gogm:"direction=outgoing;relationship=number"`
	Emails       []*Email       `gogm:"direction=outgoing;relationship=email"`
	Addresses    []*Address     `gogm:"direction=outgoing;relationship=address"`
}

type Name struct {
	gogm.BaseNode
	FullName	string `gogm:"name=limit"`
	Owner  		*Owner `gogm:"direction=incoming;relationship=name"`

}

type PhoneNumber struct {
	gogm.BaseNode
	Number  	string `gogm:"name=number"`
	Primary 	bool   
	Type    	string `gogm:"name=type"`
	Owner   	*Owner `gogm:"direction=incoming;relationship=number"`

}

type Email struct {
	gogm.BaseNode
	Address		string `gogm:"name=add"`
	Primary		bool   
	Type   		string `gogm:"name=limit"`
	Owner  		*Owner `gogm:"direction=incoming;relationship=email"`

}

type Address struct {
	gogm.BaseNode
	City       string  `gogm:"name=city"`
	Region     string  `gogm:"name=region"`
	Street     string  `gogm:"name=street"`
	PostalCode string  `gogm:"name=zip"`
	Primary    bool    
	Owner      *Owner  `gogm:"direction=incoming;relationship=address"`

}



func main() {
	config := gogm.Config{
		IndexStrategy: gogm.VALIDATE_INDEX, //other options are ASSERT_INDEX and IGNORE_INDEX
		PoolSize:      50,
		Port:          7687,
		IsCluster:     false, //tells it whether or not to use `bolt+routing`
		Host:          "0.0.0.0",
		Password:      "password",
		Username:      "neo4j",
	}

	err := gogm.Init(
		&config, 
		&Institution{}, 
		&Account{}, 
		&Balance{},
		&Owner{}, 
		&Name{}, 
		&PhoneNumber{},
		&Email{}, 
		&Address{}, 
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
	prods = 	
	Institution := &Institution{
		IntstitutionID: "woo neo4j",
		Name: "",
		Products: ,
	}

	bVal := &VertexB{
		TestTime: time.Now().UTC(),
	}


	//set bi directional pointer
	bVal.Single = aVal
	aVal.SingleA = bVal


	err = sess.SaveDepth(aVal, 3)
	if err != nil {
		panic(err)
	}



	//load the object we just made (save will set the uuid)
	var readin VertexA
	err = sess.Load(&readin, aVal.UUID)
	if err != nil {
		panic(err)
	}

	fmt.Println("done")
}
