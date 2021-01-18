package tests

import (
	"testing"

	"github.com/qweliant/neo4j/api/models"
	"github.com/stretchr/testify/assert"
)

func TestFindAllBalances(t *testing.T) {
	err := refreshNodes(sess)
	if err != nil {
		t.Errorf("this is the error getting the balance: %v\n", err)
		return
	}

	//seed db
	_, err = seedMultipleNodes(sess)
	if err != nil {
		t.Errorf("this is the error getting the balance: %v\n", err)
		return
	}

	//load the object we just made (save will set the uuid)
	var readin []models.Balance
	err = sess.LoadAll(&readin)
	if err != nil {
		t.Errorf("The error getting the balance: %v\n", err)
		return
	}

	assert.Equal(t, len(readin), 3)
	err = refreshNodes(sess)

}
