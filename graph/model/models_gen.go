// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type Center struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	PhoneNumber string        `json:"phoneNumber"`
	Memberships []*Membership `json:"memberships"`
}

type Membership struct {
	ID string `json:"id"`
	// training represents training course for the membership
	Training string    `json:"training"`
	CurrCnt  int       `json:"currCnt"`
	TotalCnt int       `json:"totalCnt"`
	Expiry   time.Time `json:"expiry"`
	// users reperesents users who share this membership
	Users []*User `json:"users"`
}

type NewUser struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
}

type User struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
}
