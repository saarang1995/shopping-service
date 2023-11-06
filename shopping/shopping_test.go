package shopping

import (
	"testing"
)


func TestAddUser(t *testing.T) {
	if len(GetAllUsers()) != 0 {
		t.Errorf("Expected UserList to be of length 0 in the beginning")
	}

	AddUser(User { Name: "User1"})

	if len(GetAllUsers()) != 1 {
		t.Errorf("UserList size is %d, expected it to be 1", len(GetAllUsers()))
	}
}


func TestGetUserById(t *testing.T) {
	createdUser := AddUser(User{Name: "User2"})

	user := GetUserById(*createdUser.Id)
	if *user.Id != *createdUser.Id {
		t.Errorf("Expected to find user with id 2")
	}


	user2 := GetUserById(3)
	if user2 != nil {
		t.Errorf("Expected to find empty user object")
	}
}