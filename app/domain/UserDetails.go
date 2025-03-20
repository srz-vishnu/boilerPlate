package domain

import "time"

type Userdetail struct {
	ID          int64     `gorm:"primaryKey"`
	Username    string    `gorm:"column:username;unique;not null"`
	Password    string    `gorm:"column:password;not null"`
	Address     string    `gorm:"column:address;not null"`
	City        string    `gorm:"column:city;not null"`
	Pincode     int64     `gorm:"column:pincode;not null"`
	Phonenumber int64     `gorm:"column:phone_number; not null"`
	Mail        string    `gorm:"column:mail;not null"`
	Status      bool      `gorm:"column:status;default:true;not null"` // Boolean field, default true, to set user active or not
	CreatedAt   time.Time `gorm:"column:created_at;autoUpdateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime"`
	UpdatedBy   *int64    `gorm:"column:deleted_by"`
}
