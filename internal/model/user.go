package model

type User struct {
	DbBasicModel
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
