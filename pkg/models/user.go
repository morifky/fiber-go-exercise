package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;not nul"`
	Email    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
}

func (u *User) FindUsers(db *gorm.DB) (*[]User, error) {
	users := []User{}

	if err := db.Debug().Model(&User{}).Limit(100).Find(&users).Error; err != nil {
		return &[]User{}, err
	}

	return &users, nil
}

func (u *User) FindUserByEmail(db *gorm.DB, email string) (*User, error) {
	user := User{}

	if err := db.Debug().Model(&User{}).Where(&User{Email: email}).Find(&user).Error; err != nil {
		return &User{}, err
	}

	return &user, nil
}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	if err := db.Debug().Create(&u).Error; err != nil {
		return &User{}, err
	}
	return u, nil
}
