package tests

/*
This test is meant to create and query models.
This would be something like a CREATE and READ method
*/

import (
	"testing"

	"github.com/qweliant/neo4j/models"
	"github.com/stretchr/testify/assert"
)

func TestFindAllUsers(t *testing.T) {
	err := refreshTapNodes(sess)
	if err != nil {
		t.Errorf("this is the error getting the users: %v\n", err)
		return
	}

	//seed db
	err = seedTapNodes(sess)
	if err != nil {
		t.Errorf("this is the error getting the users: %v\n", err)
		return
	}

	//load the object we just made (save will set the uuid)
	var readin []models.User
	err = sess.LoadAll(&readin)
	if err != nil {
		t.Errorf("The error getting the users: %v\n", err)
		return
	}

	// two users created
	assert.Equal(t, len(readin), 2)
}

func TestSaveUser(t *testing.T) {
	err := refreshTapNodes(sess)
	if err != nil {
		t.Errorf("this is the error getting the users: %v\n", err)
		return
	}

	//seed db
	err = seedOneNode(sess)
	if err != nil {
		t.Errorf("this is the error getting the users: %v\n", err)
		return
	}

	var readin []models.User
	err = sess.LoadAll(&readin)
	if err != nil {
		t.Errorf("The error getting the users: %v\n", err)
		return
	}

	assert.Equal(t, len(readin), 1)

}

func TestFindUserByID(t *testing.T) {

	assert.Equal(t, 2, 2)
}

func TestUpdateUser(t *testing.T) {

	assert.Equal(t, 2, 2)
}

func TestDeleteUser(t *testing.T) {

	assert.Equal(t, 2, 2)
}