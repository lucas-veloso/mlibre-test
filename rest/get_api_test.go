package main

import (
	"testing"
)
// el test pasa si el usuario tiene id == 1
func TestUserID(t *testing.T) {
	user, err := getApiResponse()
	if err != nil {
		t.FailNow()
	}

	if user.Id != 1 {
		t.Fail()

	}
}

// el test pasa si el usuario tiene tags
func TestTags(t *testing.T) {
	user, err := getApiResponse()
	if err != nil {
		t.FailNow()
	}

	if len(user.Tags) == 0 {
		t.Fail()
	}
}


