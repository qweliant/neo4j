package main

import (
	"fmt"

	"github.com/mindstand/gogm"
)

type User struct {
	gogm.BaseNode

	UserID string  `gogm:"name=user_id"`
	Items  []*Item `gogm:"direction=outgoing;relationship=item"`
}

type Item struct {
	gogm.BaseNode

	User        *User        `gogm:"direction=incoming;relationship=item"`
	Institution *Institution `gogm:"direction=outgoing;relationship=institution"`
	Account     []*Account   `gogm:"direction=outgoing;relationship=account"`
}

type Institution struct {
	// provides required node fields
	gogm.BaseNode

	IntstitutionID string `gogm:"name=inst_id"`
	Name           string `gogm:"name=name"`
	Products       []string
	PrimaryColor   string `gogm:"name=prim_color"`
	Logo           string `gogm:"name=logo"`
	Item           *Item  `gogm:"direction=incoming;relationship=institution"`
}

type Account struct {
	gogm.BaseNode

	AccountID          string `gogm:"name=accnt_id"`
	Name               string `gogm:"name=name"`
	OfficialName       string `gogm:"name=offic_name"`
	Type               string `gogm:"name=type"`
	Subtype            string `gogm:"name=subtype"`
	IntstitutionID     string `gogm:"name=institution_id"`
	VerificationStatus bool
	Owner              *Owner   `gogm:"direction=outgoing;relationship=owner"`
	Balance            *Balance `gogm:"direction=outgoing;relationship=balance"`
	Item               *Item    `gogm:"direction=incoming;relationship=account"`
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

	Account      *Account       `gogm:"direction=incoming;relationship=owner"`
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
		&User{},
		&Item{},
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
		AccountID:          "3",
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

	// fill in data
	n := &Name{FullName: "Monique"}
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

	p3 := &PhoneNumber{
		Number:  "877-226-7723",
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

	e3 := &Email{
		Address: "hawtpower@foucault.com",
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
		Region:     "DC",
		Street:     "123 Chocolate City Ave",
		PostalCode: "12738",
		Primary:    false,
	}

	a3 := &Address{
		City:       "Dallas",
		Region:     "TX",
		Street:     "123 Aye Baybay",
		PostalCode: "77232",
		Primary:    false,
	}

	// declare array for owners
	var accountNames []*Name
	var accountPhoneNumber []*PhoneNumber
	var accountEmail []*Email
	var accountAddress []*Address

	// append together for owners
	phoneNums := append(accountPhoneNumber, p1)
	names := append(accountNames, n1, n)
	emails := append(accountEmail, e1)
	addresses := append(accountAddress, a1)

	// declare array for owners
	var accountNames2 []*Name
	var accountPhoneNumber2 []*PhoneNumber
	var accountEmail2 []*Email
	var accountAddress2 []*Address

	phoneNums2 := append(accountPhoneNumber2, p2)
	names2 := append(accountNames2, n2)
	emails2 := append(accountEmail2, e2)
	addresses2 := append(accountAddress2, a2)

	// declare array for owners
	var accountNames3 []*Name
	var accountPhoneNumber3 []*PhoneNumber
	var accountEmail3 []*Email
	var accountAddress3 []*Address

	phoneNums3 := append(accountPhoneNumber3, p3)
	names3 := append(accountNames3, n3)
	emails3 := append(accountEmail3, e3)
	addresses3 := append(accountAddress3, a3)

	owners1 := &Owner{
		Names:        names,
		PhoneNumbers: phoneNums,
		Emails:       emails,
		Addresses:    addresses,
	}

	owners2 := &Owner{
		Names:        names2,
		PhoneNumbers: phoneNums2,
		Emails:       emails2,
		Addresses:    addresses2,
	}
	owners3 := &Owner{
		Names:        names3,
		PhoneNumbers: phoneNums3,
		Emails:       emails3,
		Addresses:    addresses3,
	}

	// set bi directional pointer
	accountA.Balance = balance1
	balance1.Account = accountA

	accountB.Balance = balance2
	balance2.Account = accountB

	accountC.Balance = balance3
	balance3.Account = accountC

	owners1.Account = accountA
	accountA.Owner = owners1

	owners2.Account = accountB
	accountB.Owner = owners2

	owners3.Account = accountC
	accountC.Owner = owners3

	user1 := &User{
		UserID: "1",
	}

	user2 := &User{
		UserID: "2",
	}

	i1 := &Item{}

	i2 := &Item{}


	var listOfAccnt1 []*Account
	var listOfAccnt2 []*Account

	accnts1 := append(listOfAccnt1, accountA, accountB)
	accnts2 := append(listOfAccnt2, accountC)

	i1.Account = accnts1
	i1.Institution = institutionA

	i2.Account = accnts2
	i2.Institution = institutionB
	
	var listOfItems1 []*Item
	var listOfItems2 []*Item
	
	item1 := append(listOfItems1, i1)
	item2 := append(listOfItems2, i2)

	user1.Items = item1
	user2.Items = item2


	err = sess.SaveDepth(user1, 8)
	if err != nil {
		panic(err)
	}
	err = sess.SaveDepth(user2, 8)
	if err != nil {
		panic(err)
	}
	//load the object we just made (save will set the uuid)
	var readin User
	err = sess.Load(&readin, user1.UUID)
	if err != nil {
		panic(err)
	}
	//load the object we just made (save will set the uuid)
	var readin2 User
	err = sess.Load(&readin2, user2.UUID)
	if err != nil {
		panic(err)
	}
	fmt.Println("Done")
}
