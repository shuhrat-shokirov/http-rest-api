package model

import "testing"

func TestUser(t *testing.T) *User {
	return &User{
		Email:    "abcd@mail.ru",
		Password: "password",
	}
}
