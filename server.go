package main

import (
	"fmt"

	"github.com/mindstand/gogm"
)

type tdString string
type tdInt int

type Institution struct {
	// provides required node fields
	gogm.BaseNode
	IntstitutionID string `gogm:"name=inst_id"`
	Name           string `gogm:"name=name"`
	Products       []string
	PrimaryColor   string     `gogm:"name=prim_color"`
	Logo           string     `gogm:"name=logo"`
	Accounts       []*Account `gogm:"direction=outgoing;relationship=account"`
}

type Account struct {
	gogm.BaseNode
	AccountID          string `gogm:"name=accnt_id"`
	Name               string `gogm:"name=name"`
	OfficialName       string `gogm:"name=offic_name"`
	Type               string `gogm:"name=type"`
	Subtype            string `gogm:"name=subtype"`
	VerificationStatus bool
	Owner              *Owner       `gogm:"direction=outgoing;relationship=owner"`
	Balance            *Balance     `gogm:"direction=outgoing;relationship=balance"`
	Institution        *Institution `gogm:"direction=incoming;relationship=account"`
}

type Balance struct {
	gogm.BaseNode
	Available float64  `gogm:"name=avail"`
	Current   float64  `gogm:"name=current"`
	Limit     float64  `gogm:"name=limit"`
	Currency  string   `gogm:"name=currency"`
	Account   *Account `gogm:"direction=incoming;relationship=balance"`
}

type Owner struct {
	gogm.BaseNode
	Accounts     *Account       `gogm:"direction=incoming;relationship=owner"`
	Names        []*Name        `gogm:"direction=outgoing;relationship=name"`
	PhoneNumbers []*PhoneNumber `gogm:"direction=outgoing;relationship=number"`
	Emails       []*Email       `gogm:"direction=outgoing;relationship=email"`
	Addresses    []*Address     `gogm:"direction=outgoing;relationship=address"`
}

type Name struct {
	gogm.BaseNode
	FullName string `gogm:"name=limit"`
	Owner    *Owner `gogm:"direction=incoming;relationship=name"`
}

type PhoneNumber struct {
	gogm.BaseNode
	Number  string `gogm:"name=number"`
	Primary bool
	Type    string `gogm:"name=type"`
	Owner   *Owner `gogm:"direction=incoming;relationship=number"`
}

type Email struct {
	gogm.BaseNode
	Address string `gogm:"name=add"`
	Primary bool
	Type    string `gogm:"name=limit"`
	Owner   *Owner `gogm:"direction=incoming;relationship=email"`
}

type Address struct {
	gogm.BaseNode
	City       string `gogm:"name=city"`
	Region     string `gogm:"name=region"`
	Street     string `gogm:"name=street"`
	PostalCode string `gogm:"name=zip"`
	Primary    bool
	Owner      *Owner `gogm:"direction=incoming;relationship=address"`
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

	// Create onjects to graph
	var prods = []string{"auth", "tranactions", "identity"}
	institutionA := &Institution{
		IntstitutionID: "1",
		Name:           "Bank of Dees Nutz Nigguh",
		Products:       prods,
		PrimaryColor:   "red",
		Logo:           "A image byte string was here",
	}

	institutionB := &Institution{
		IntstitutionID: "2",
		Name:           "Bank of Black Excellence",
		Products:       prods,
		PrimaryColor:   "black",
		Logo:           "A image byte string was here",
	}

	accountA := &Account{
		AccountID:          "1",
		Name:               "BODNN Checking",
		OfficialName:       "Checking Account",
		Type:               "depository",
		Subtype:            "checking",
		VerificationStatus: true,
	}

	accountB := &Account{
		AccountID:          "2",
		Name:               "BODNN Savings",
		OfficialName:       "Savings Account",
		Type:               "depository",
		Subtype:            "savings",
		VerificationStatus: true,
	}

	accountC := &Account{
		AccountID:          "1",
		Name:               "BBE Savings",
		OfficialName:       "BBE Savings Account",
		Type:               "depository",
		Subtype:            "savings",
		VerificationStatus: true,
	}

	balance1 := &Balance{
		Available: 110,
		Current:   110,
		Limit:     0,
		Currency:  "USD",
	}
	balance2 := &Balance{
		Available: 110,
		Current:   110,
		Limit:     0,
		Currency:  "USD",
	}
	balance3 := &Balance{
		Available: 110,
		Current:   110,
		Limit:     0,
		Currency:  "USD",
	}
	// declare array for owners
	var accountNames []*Name
	var accountPhoneNumber []*PhoneNumber
	var accountEmail []*Email
	var accountAddress []*Address

	// array for instituion accounts
	var accounts []*Account

	// fill in data
	n1 := &Name{FullName: "Latasha"}
	n2 := &Name{FullName: "Derrick"}
	n3 := &Name{FullName: "Aldean"}

	p1 := &PhoneNumber{
		Number:  "961-867-5309",
		Primary: false,
		Type:    "cell",
	}

	p2 := &PhoneNumber{
		Number:  "281-330-8004",
		Primary: false,
		Type:    "home",
	}

	e1 := &Email{
		Address: "ddennat@qualia.com",
		Primary: false,
		Type:    "primary",
	}

	e2 := &Email{
		Address: "loveurself@gmail.com",
		Primary: false,
		Type:    "secondary",
	}

	a1 := &Address{
		City:       "Atlanta",
		Region:     "GA",
		Street:     "1017 Brick Rd, 745",
		PostalCode: "39827",
		Primary:    true,
	}

	a2 := &Address{
		City:       "Washington DC",
		Region:     "District of Columbia",
		Street:     "123 Chocolate City Ave",
		PostalCode: "12738",
		Primary:    false,
	}

	// append together for owners
	phoneNums := append(accountPhoneNumber, p1, p2)
	names := append(accountNames, n1, n2, n3)
	emails := append(accountEmail, e1, e2)
	addresses := append(accountAddress, a1, a2)

	owners1 := &Owner{
		Names:        names,
		PhoneNumbers: phoneNums,
		Emails:       emails,
		Addresses:    addresses,
	}

	owners2 := &Owner{
		Names:        names,
		PhoneNumbers: phoneNums,
		Emails:       emails,
		Addresses:    addresses,
	}
	owners3 := &Owner{
		Names:        names,
		PhoneNumbers: phoneNums,
		Emails:       emails,
		Addresses:    addresses,
	}
	// set bi directional pointer
	accountA.Balance = balance1
	balance1.Account = accountA

	accountB.Balance = balance2
	balance2.Account = accountB

	accountC.Balance = balance3
	balance3.Account = accountC

	accountA.Owner = owners1
	owners1.Accounts = accountA

	accountB.Owner = owners2
	owners2.Accounts = accountB

	accountC.Owner = owners3
	owners3.Accounts = accountC

	accountsInsA := append(accounts, accountA, accountB)
	accountsInsB := append(accounts, accountB, accountC)

	institutionA.Accounts = accountsInsA
	institutionB.Accounts = accountsInsB

	err = sess.SaveDepth(institutionA, 2)
	if err != nil {
		panic(err)
	}
	err = sess.SaveDepth(institutionB, 2)
	if err != nil {
		panic(err)
	}
	//load the object we just made (save will set the uuid)
	var readin Institution
	err = sess.Load(&readin, institutionA.UUID)
	if err != nil {
		panic(err)
	}
	//load the object we just made (save will set the uuid)
	var readin2 Institution
	err = sess.Load(&readin2, institutionB.UUID)
	if err != nil {
		panic(err)
	}
	fmt.Println("Done")
}
