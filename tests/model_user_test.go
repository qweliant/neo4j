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
	_, err = seedOneNode(sess)
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
	err := refreshTapNodes(sess)
	if err != nil {
		t.Errorf("this is the error getting the users: %v\n", err)
		return
	}

	//seed db
	id, err := seedOneNode(sess)
	if err != nil {
		t.Errorf("this is the error getting the users: %v\n", err)
		return
	}

	var readin models.User
	err = sess.Load(&readin, id)
	if err != nil {
		t.Errorf("The error getting the users: %v\n", err)
		return
	}
	assert.Equal(t, readin.UUID, id)
}

func TestUpdateUser(t *testing.T) {
	err := refreshTapNodes(sess)
	if err != nil {
		t.Errorf("this is the error getting the users: %v\n", err)
		return
	}

	//seed db
	id, err := seedOneNode(sess)
	if err != nil {
		t.Errorf("this is the error getting the users: %v\n", err)
		return
	}

	var readin models.User
	err = sess.Load(&readin, id)
	if err != nil {
		t.Errorf("The error getting the users: %v\n", err)
		return
	}
	assert.Equal(t, readin.UserID, "1")

	var copy models.User
	copy = readin

	readin.UserID = "3"

	err = sess.Save(&readin)
	if err != nil {
		panic(err)
	}

	var readBackIn models.User
	err = sess.Load(&readBackIn, id)
	if err != nil {
		t.Errorf("The error getting the users: %v\n", err)
		return
	}
	assert.NotEqual(t, readBackIn.UserID, copy.UserID)

}

func TestDeleteUser(t *testing.T) {
	err := refreshTapNodes(sess)
	if err != nil {
		t.Errorf("this is the error getting the users: %v\n", err)
		return
	}

	//seed db
	_, err = seedOneNode(sess)
	if err != nil {
		t.Errorf("this is the error getting the users: %v\n", err)
		return
	}

	var readin models.User
	err = sess.Delete(&readin)
	if err != nil {
		t.Errorf("The error delting the user: %v\n", err)
		return
	}
	
	err = sess.DeleteUUID(readin.UUID)
	if err != nil {
		t.Errorf("The error getting the users: %v\n", err)
		return
	}
}
