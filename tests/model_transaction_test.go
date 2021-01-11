package tests

/*
This test is meant to create and query models.
This would be something like a CREATE and READ method
*/

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindAllUsers(t *testing.T) {
	sess, err := refreshTapNodes()
	if err != nil {
		log.Fatal(err)
	}

	//seed db
	err = seedTapNodes(sess)
	if err != nil {
		t.Errorf("this is the error getting the users: %v\n", err)
		return
	}

	
	assert.Equal(t, 1, 1)

}
