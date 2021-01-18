package tests

/*
This test is meant to create and query models.
This would be something like a CREATE and READ method
*/

import (
	"fmt"
	"testing"

	"github.com/qweliant/neo4j/api/models"
	"github.com/stretchr/testify/assert"
)

func TestFindAllUsers(t *testing.T) {
	err := refreshNodes(sess)
	if err != nil {
		t.Errorf("this is the error getting the users: %v\n", err)
		return
	}

	//seed db
	_, err = seedMultipleNodes(sess)
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
	err = refreshNodes(sess)

}

func TestSaveUser(t *testing.T) {
	err := refreshNodes(sess)
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
	err = refreshNodes(sess)

}

func TestFindUserByID(t *testing.T) {
	err := refreshNodes(sess)
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
	err = refreshNodes(sess)

}

func TestUpdateUser(t *testing.T) {
	err := refreshNodes(sess)
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
	assert.Equal(t, readin.UserID, int64(1))

	var copy models.User
	copy = readin

	readin.UserID = int64(3)

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
	err = refreshNodes(sess)

}

func TestDeleteUser(t *testing.T) {
	err := refreshNodes(sess)
	if err != nil {
		t.Errorf("this is the error getting the users: %v\n", err)
		return
	}

	//seed db
	id, err := seedMultipleNodes(sess)
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

	// will get an error that kicks out of test because nothing returned
	query := fmt.Sprintf("MATCH (n:User {uuid: '%s'})-[*0..]->(x) DETACH DELETE x", id)
	_ = sess.Query(query, nil, readin)

	// doesnt ret anything because of deletion
	var readBackIn models.User
	_ = sess.Load(&readBackIn, id)

	assert.NotEqual(t, readBackIn.UUID, id)
	err = refreshNodes(sess)

}
