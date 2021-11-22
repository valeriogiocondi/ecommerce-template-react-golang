package customer

import (
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	Id         uint64    `gorm:"column:id; autoIncrement"`
	UUID       uuid.UUID `gorm:"column:uuid; type:uuid; primaryKey"`
	FirebaseId string    `gorm:"column:firebase_id; not null"`
	FirstName  string    `gorm:"column:first_name; not null"`
	LastName   string    `gorm:"column:last_name; not null"`
	Email      string    `gorm:"column:email; unique; not null"`
	Tel        string    `gorm:"column:tel; not null"`
	Address    string    `gorm:"column:address; not null"`
	Num        string    `gorm:"column:num; not null"`
	Cap        string    `gorm:"column:cap; not null"`
	City       string    `gorm:"column:city; not null"`
	State      string    `gorm:"column:state; not null"`
	CreatedAt  time.Time `gorm:"column:date_create"`
}
