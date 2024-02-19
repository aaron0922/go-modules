package user

import (
	"errors"
)

type User struct {
	ID int
	Name string
}

var users = map[int]User{
	1: User{ID: 1, Name: "User 1"},
	2: User{ID: 2, Name: "User 2"},
	3: User{ID: 3, Name: "User 3"},
}


func GetUserByID(id int) (*User, error) {
	if id > 3 {
		return nil, errors.New("User not found")
	}
	
	user := users[id]
	return &user, nil
}
