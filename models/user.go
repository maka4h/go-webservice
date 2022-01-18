package models

import (
	"errors"
	"fmt"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("new user must not include ID")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

func GetUserByID(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}
	return User{}, fmt.Errorf("user with ID '%v' not found", id)
}

func UpdateUserByID(u User) (User, error) {
	for i, candidate := range users {
		if u.ID == candidate.ID {
			users[i] = &u
			fmt.Println("User has been updated", u.FirstName)
			return *users[i], nil
		}
	}
	return User{}, fmt.Errorf("user with ID '%v' not found", u.ID)
}

func RemoveUserByID(id int) (User, error) {
	for i, u := range users {
		if u.ID == id {
			toDelete, _ := GetUserByID(id)
			fmt.Println("Going to remove user ", u.ID)
			users = append(users[:i], users[i+1:]...)
			return toDelete, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", id)
}
