package tests

import (
	"fmt"
	"testing"

	"github.com/qweliant/neo4j/api/models"
	"github.com/stretchr/testify/assert"
)

func TestFindAllTransactions(t *testing.T) {
	err := refreshNodes(sess)
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
	err = refreshNodes(sess)

}

// gets transactions for account id
func TestFindUserTransaction(t *testing.T) {
	err := refreshNodes(sess)
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
	err = refreshNodes(sess)

}

func TestAddUserTransaction(t *testing.T) {
	err := refreshNodes(sess)
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

	assert.Equal(t, len(readin), 3)

	transaction6 := &models.Transaction{
		Name:           "Netflix Subscription",
		MerchantName:   "Netflix",
		Ammount:        10.13,
		Currency:       "USD",
		PaymentChannel: "online",
		Pending:        false,
	}

	readin = append(readin, *transaction6)

	assert.Equal(t, len(readin), 4)
	err = refreshNodes(sess)

}
