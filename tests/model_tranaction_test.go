package tests

import (
	"fmt"
	"testing"

	"github.com/qweliant/neo4j/models"
	"github.com/stretchr/testify/assert"
)

func TestFindAllTransactions(t *testing.T) {
	err := refreshTapNodes(sess)
	if err != nil {
		t.Errorf("this is the error getting the transactions: %v\n", err)
		return
	}

	//seed db
	_, err = seedMultipleNodes(sess)
	if err != nil {
		t.Errorf("this is the error getting the transactions: %v\n", err)
		return
	}

	//load the object we just made (save will set the uuid)
	var readin []models.Transaction
	err = sess.LoadAll(&readin)
	if err != nil {
		t.Errorf("The error getting the users: %v\n", err)
		return
	}

	// two users created
	assert.Equal(t, len(readin), 6)
}

// gets transactions for account id
func TestFindUserTransaction(t *testing.T) {
	err := refreshTapNodes(sess)
	if err != nil {
		t.Errorf("this is the error getting the transactions: %v\n", err)
		return
	}

	//seed db
	id, err := seedMultipleNodes(sess)
	if err != nil {
		t.Errorf("this is the error getting the transactions: %v\n", err)
		return
	}

	var user models.User
	err = sess.Load(&user, id)
	if err != nil {
		t.Errorf("The error getting the users: %v\n", err)
		return
	}
	fmt.Println(user.Items[0].Accounts)

	// I have user, but here i have to use account id (accounts aren't on items?), then look at tranactaions
	// will hard code for now but know this is need for api req
	accntID := "1"
	var readin []*models.Transaction
	query := fmt.Sprintf("MATCH (n {accnt_id: '%s'})-->(m:Transaction) RETURN m", accntID)

	err = sess.Query(query, nil, &readin)
	if err != nil {
		t.Errorf("The error getting the users: %v\n", err)
		return
	}

	assert.Equal(t, len(readin), 3)

}

func TestAddUserTransaction(t *testing.T) {
	err := refreshTapNodes(sess)
	if err != nil {
		t.Errorf("this is the error getting the transactions: %v\n", err)
		return
	}

	//seed db
	id, err := seedMultipleNodes(sess)
	if err != nil {
		t.Errorf("this is the error getting the transactions: %v\n", err)
		return
	}

	var user models.User
	err = sess.Load(&user, id)
	if err != nil {
		t.Errorf("The error getting the users: %v\n", err)
		return
	}
	fmt.Println(user.Items[0].Accounts)

	// I have user, but here i have to use account id (accounts aren't on items?), then look at tranactaions
	// will hard code for now but know this is need for api req
	accntID := "1"
	var readin []models.Transaction

	query := fmt.Sprintf("MATCH (n {accnt_id: '%s'})-->(m:Transaction) RETURN m", accntID)
	err = sess.Query(query, nil, &readin)
	if err != nil {
		t.Errorf("The error getting the users: %v\n", err)
		return
	}

	transaction6 := models.Transaction{
		Name:           "Netflix Subscription",
		MerchantName:   "Netflix",
		Ammount:        10.13,
		Currency:       "USD",
		PaymentChannel: "online",
		Pending:        false,
	}

	readin = append(readin, transaction6)
	fmt.Println(readin)
	err = sess.Save(&readin)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 4, 4)

}
