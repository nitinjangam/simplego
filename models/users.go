package models

import "os"

//Users struct
type Users struct {
	UserList []User `json:"users"`
}

//User struct
type User struct {
	ID     int    `json:"userid"`
	Name   string `json:"name"`
	Active string `json:"active"`
}

func (u *Users) getAllUsers(f *os.File) error {
	return nil
}

func (u *User) getUserData(f *os.File) error {
	return nil
}

func (u *User) updateUserData(f *os.File) error {
	return nil
}

func (u *User) deleteUser(f *os.File) error {
	return nil
}
