package entity

import (
	"gorm.io/gorm"
)

type User struct{
	gorm.Model
	StudentID string `valid:"required~StudentID is required, matches(^[MDB]\\d{7}$)"`
	FirstName string `valid:"required~FirstName is required"`
}