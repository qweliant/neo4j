package tests

import (
	"fmt"

	"github.com/mindstand/gogm"
	"github.com/qweliant/neo4j/api/models"
)

func Database() (*gogm.Session, error) {
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
	return sess, nil
}

var (
	sess, err = Database()
)

func refreshNodes(s *gogm.Session) error {

	err = s.PurgeDatabase()
	if err != nil {
		panic(err)
	}
	return nil
}

func seedMultipleNodes(sess *gogm.Session) (string, error) {
	var prods = []string{"auth", "tranactions", "identity"}
	var cats = []string{"Shops", "Computers and Electronics"}

	institutionA := &models.Institution{
		IntstitutionID: "1",
		Name:           "Bank of Dees",
		Products:       prods,
		PrimaryColor:   "red",
		Logo:           "A image byte string was here",
	}

	institutionB := &models.Institution{
		IntstitutionID: "2",
		Name:           "Bank of Black Excellence",
		Products:       prods,
		PrimaryColor:   "black",
		Logo:           "A image byte string was here",
	}

	accountA := &models.Account{
		AccountID:          "1",
		Name:               "BODNN Checking",
		OfficialName:       "Checking Account",
		Type:               "depository",
		Subtype:            "checking",
		VerificationStatus: true,
	}

	accountB := &models.Account{
		AccountID:          "2",
		Name:               "BODNN Savings",
		OfficialName:       "Savings Account",
		Type:               "depository",
		Subtype:            "savings",
		VerificationStatus: true,
	}

	accountC := &models.Account{
		AccountID:          "3",
		Name:               "BBE Savings",
		OfficialName:       "BBE Savings Account",
		Type:               "depository",
		Subtype:            "savings",
		VerificationStatus: true,
	}

	balance1 := &models.Balance{
		Available: 110,
		Current:   110,
		Limit:     4345,
		Currency:  "USD",
	}
	balance2 := &models.Balance{
		Available: 300,
		Current:   300,
		Limit:     0,
		Currency:  "USD",
	}
	balance3 := &models.Balance{
		Available: 12,
		Current:   45,
		Limit:     0,
		Currency:  "USD",
	}

	location := &models.Location{
		Address:     "123 Forest Road",
		City:        "Seattle",
		Region:      "WA",
		PostalCode:  "12345",
		Country:     "US",
		StoreNumber: "1235",
		Lat:         40.740352,
		Lon:         -74.001761,
	}

	location2 := &models.Location{
		Address:     "123 Forest Road",
		City:        "Seattle",
		Region:      "WA",
		PostalCode:  "12345",
		Country:     "US",
		StoreNumber: "1235",
		Lat:         40.740352,
		Lon:         -74.001761,
	}

	transaction := &models.Transaction{
		Name:           "Apple Store",
		MerchantName:   "Apple",
		Ammount:        2307.23,
		Currency:       "USD",
		Category:       cats,
		PaymentChannel: "in store",
		Pending:        false,
		Location:       location,
		Account:        accountA,
	}

	transaction2 := &models.Transaction{
		Name:           "Google Store",
		MerchantName:   "Google",
		Ammount:        207.47,
		Currency:       "USD",
		Category:       cats,
		PaymentChannel: "in store",
		Pending:        false,
		Location:       location,
		Account:        accountA,
	}

	transaction3 := &models.Transaction{
		Name:           "Amazon Store",
		MerchantName:   "Amazon",
		Ammount:        42.83,
		Currency:       "USD",
		Category:       cats,
		PaymentChannel: "in store",
		Pending:        false,
		Location:       location,
		Account:        accountA,
	}

	transaction4 := &models.Transaction{
		Name:           "Apple Store",
		MerchantName:   "Apple",
		Ammount:        2307.23,
		Currency:       "USD",
		Category:       cats,
		PaymentChannel: "in store",
		Pending:        false,
		Location:       location2,
		Account:        accountB,
	}

	transaction5 := &models.Transaction{
		Name:           "Google Store",
		MerchantName:   "Google",
		Ammount:        207.47,
		Currency:       "USD",
		Category:       cats,
		PaymentChannel: "in store",
		Pending:        false,
		Location:       location2,
		Account:        accountB,
	}

	transaction6 := &models.Transaction{
		Name:           "Amazon Store",
		MerchantName:   "Amazon",
		Ammount:        42.83,
		Currency:       "USD",
		Category:       cats,
		PaymentChannel: "in store",
		Pending:        false,
		Location:       location2,
		Account:        accountB,
	}

	var accountTransactions []*models.Transaction
	accntTransactions := append(accountTransactions, transaction, transaction2, transaction3)

	var accountTransactions2 []*models.Transaction
	accntTransactions2 := append(accountTransactions2, transaction4, transaction5, transaction6)

	// fill in data
	n := &models.Name{FullName: "Monique"}
	n1 := &models.Name{FullName: "Latasha"}
	n2 := &models.Name{FullName: "Derrick"}
	n3 := &models.Name{FullName: "Aldean"}

	p1 := &models.PhoneNumber{
		Number:  "961-867-5309",
		Primary: false,
		Type:    "cell",
	}

	p2 := &models.PhoneNumber{
		Number:  "281-330-8004",
		Primary: false,
		Type:    "home",
	}

	p3 := &models.PhoneNumber{
		Number:  "877-226-7723",
		Primary: false,
		Type:    "home",
	}
	e1 := &models.Email{
		Address: "ddennat@qualia.com",
		Primary: false,
		Type:    "primary",
	}

	e2 := &models.Email{
		Address: "loveurself@gmail.com",
		Primary: false,
		Type:    "secondary",
	}

	e3 := &models.Email{
		Address: "hawtpower@foucault.com",
		Primary: false,
		Type:    "secondary",
	}

	a1 := &models.Address{
		City:       "Atlanta",
		Region:     "GA",
		Street:     "1017 Brick Rd, 745",
		PostalCode: "39827",
		Primary:    true,
	}

	a2 := &models.Address{
		City:       "Washington DC",
		Region:     "DC",
		Street:     "123 Chocolate City Ave",
		PostalCode: "12738",
		Primary:    false,
	}

	a3 := &models.Address{
		City:       "Dallas",
		Region:     "TX",
		Street:     "123 Aye Baybay",
		PostalCode: "77232",
		Primary:    false,
	}

	// declare array for owners
	var accountNames []*models.Name
	var accountPhoneNumber []*models.PhoneNumber
	var accountEmail []*models.Email
	var accountAddress []*models.Address

	// append together for owners
	phoneNums := append(accountPhoneNumber, p1)
	names := append(accountNames, n1, n)
	emails := append(accountEmail, e1)
	addresses := append(accountAddress, a1)

	// declare array for owners
	var accountNames2 []*models.Name
	var accountPhoneNumber2 []*models.PhoneNumber
	var accountEmail2 []*models.Email
	var accountAddress2 []*models.Address

	phoneNums2 := append(accountPhoneNumber2, p2)
	names2 := append(accountNames2, n2)
	emails2 := append(accountEmail2, e2)
	addresses2 := append(accountAddress2, a2)

	// declare array for owners
	var accountNames3 []*models.Name
	var accountPhoneNumber3 []*models.PhoneNumber
	var accountEmail3 []*models.Email
	var accountAddress3 []*models.Address

	phoneNums3 := append(accountPhoneNumber3, p3)
	names3 := append(accountNames3, n3)
	emails3 := append(accountEmail3, e3)
	addresses3 := append(accountAddress3, a3)

	owners1 := &models.Owner{
		Names:        names,
		PhoneNumbers: phoneNums,
		Emails:       emails,
		Addresses:    addresses,
	}

	owners2 := &models.Owner{
		Names:        names2,
		PhoneNumbers: phoneNums2,
		Emails:       emails2,
		Addresses:    addresses2,
	}
	owners3 := &models.Owner{
		Names:        names3,
		PhoneNumbers: phoneNums3,
		Emails:       emails3,
		Addresses:    addresses3,
	}

	// set bi directional pointer
	accountA.Balance = balance1
	balance1.Account = accountA
	accountA.Transactions = accntTransactions

	accountB.Balance = balance2
	balance2.Account = accountB
	accountB.Transactions = accntTransactions2

	accountC.Balance = balance3
	balance3.Account = accountC

	owners1.Account = accountA
	accountA.Owner = owners1

	owners2.Account = accountB
	accountB.Owner = owners2

	owners3.Account = accountC
	accountC.Owner = owners3

	user1 := &models.User{
		UserID: 1,
	}

	user2 := &models.User{
		UserID: 2,
	}

	var listOfAccnt1 []*models.Account
	var listOfAccnt2 []*models.Account

	accnts1 := append(listOfAccnt1, accountA, accountB)
	accnts2 := append(listOfAccnt2, accountC)

	i1 := &models.Item{
		Institution: institutionA,
	}
	i2 := &models.Item{
		Institution: institutionB,
	}
	accountA.Item = i1

	accountB.Item = i1
	i1.Accounts = accnts1

	accountC.Item = i2
	i2.Accounts = accnts2

	var listOfItems1 []*models.Item
	var listOfItems2 []*models.Item

	item1 := append(listOfItems1, i1)
	item2 := append(listOfItems2, i2)

	user1.Items = item1
	user2.Items = item2

	err := sess.SaveDepth(user1, 10)
	if err != nil {
		panic(err)
	}
	err = sess.SaveDepth(user2, 10)
	if err != nil {
		panic(err)
	}

	fmt.Println("Done Seeding")
	return user2.UUID, nil
}

func seedOneNode(sess *gogm.Session) (string, error) {
	var prods = []string{"auth", "tranactions", "identity"}

	institutionB := &models.Institution{
		IntstitutionID: "2",
		Name:           "Bank of Black Excellence",
		Products:       prods,
		PrimaryColor:   "black",
		Logo:           "A image byte string was here",
	}

	accountC := &models.Account{
		AccountID:          "3",
		Name:               "BBE Savings",
		OfficialName:       "BBE Savings Account",
		Type:               "depository",
		Subtype:            "savings",
		VerificationStatus: true,
	}

	balance3 := &models.Balance{
		Available: 110,
		Current:   110,
		Limit:     0,
		Currency:  "USD",
	}

	// fill in data
	n3 := &models.Name{FullName: "Aldean"}

	p3 := &models.PhoneNumber{
		Number:  "877-226-7723",
		Primary: false,
		Type:    "home",
	}

	e3 := &models.Email{
		Address: "hawtpower@foucault.com",
		Primary: false,
		Type:    "secondary",
	}

	a3 := &models.Address{
		City:       "Dallas",
		Region:     "TX",
		Street:     "123 Aye Baybay",
		PostalCode: "77232",
		Primary:    false,
	}

	// declare array for owners
	var accountNames3 []*models.Name
	var accountPhoneNumber3 []*models.PhoneNumber
	var accountEmail3 []*models.Email
	var accountAddress3 []*models.Address

	phoneNums3 := append(accountPhoneNumber3, p3)
	names3 := append(accountNames3, n3)
	emails3 := append(accountEmail3, e3)
	addresses3 := append(accountAddress3, a3)

	owners3 := &models.Owner{
		Names:        names3,
		PhoneNumbers: phoneNums3,
		Emails:       emails3,
		Addresses:    addresses3,
	}

	// set bi directional pointer

	accountC.Balance = balance3
	balance3.Account = accountC

	owners3.Account = accountC
	accountC.Owner = owners3

	user2 := &models.User{
		UserID: 1,
	}

	i2 := &models.Item{}

	var listOfAccnt2 []*models.Account

	accnts2 := append(listOfAccnt2, accountC)

	i2.Accounts = accnts2
	i2.Institution = institutionB

	var listOfItems2 []*models.Item

	item2 := append(listOfItems2, i2)

	user2.Items = item2

	err = sess.SaveDepth(user2, 8)
	if err != nil {
		panic(err)
	}

	fmt.Println("Done Seeding")
	return user2.UUID, nil
}
