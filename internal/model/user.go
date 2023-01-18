package model

import (
	"time"
)

type User struct {
	Id          string     `json:"id" bson:"id"`
	Username    string     `json: "username" bson:"username"`
	Email       string     `json:"email" bson:"email"`
	Phone       string     `json:"phone" bson:"phone"`
	DateOfBirth *time.Time `json: "dateOfBirth" bson:"dateOfBirth"`
}
