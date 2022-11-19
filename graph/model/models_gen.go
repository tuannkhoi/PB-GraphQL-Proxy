// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewUser struct {
	Username    string  `json:"username"`
	Password    string  `json:"password"`
	Email       *string `json:"email"`
	PhoneNumber *string `json:"phoneNumber"`
}

type User struct {
	Username      string  `json:"username"`
	Password      string  `json:"password"`
	Email         *string `json:"email"`
	EmailVerified *bool   `json:"emailVerified"`
	PhoneNumber   *string `json:"phoneNumber"`
}
