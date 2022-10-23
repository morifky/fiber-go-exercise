package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint32 `gorm:"primary_key;auto_increment"`
	Username string `gorm:"not null"`
	Email    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
}

type UserRepository struct {
	db *gorm.DB
}

func (ur *UserRepository) FindAll() (*[]User, error) {
	users := []User{}

	if err := ur.db.Debug().Model(&User{}).Limit(100).Find(&users).Error; err != nil {
		return &[]User{}, err
	}

	return &users, nil
}

func (ur *UserRepository) FindByEmail(email string) (*User, error) {
	user := User{}

	if err := ur.db.Debug().Model(&User{}).Where("email = ?", email).Take(&user).Error; err != nil {
		return &User{}, err
	}

	return &user, nil
}

func (ur *UserRepository) FindOne(id uint32) (*User, error) {
	user := User{}

	if err := ur.db.Debug().Model(&User{}).Where("id = ?", id).Take(&user).Error; err != nil {
		return &User{}, err
	}

	return &user, nil
}

func (ur *UserRepository) UpdateOne(id uint32, data *User) error {

	db := ur.db.Debug().Model(&User{}).Where("id = ?", id).Take(&User{}).Updates(data)

	if db.Error != nil {
		return db.Error
	}
	return nil

}

func (ur *UserRepository) DeleteOne(id uint32) (int64, error) {
	db := ur.db.Debug().Model(&User{}).Where("id = ?", id).Take(&User{}).Delete(&User{})
	if db.Error != nil {
		return 0, db.Error
	}

	return db.RowsAffected, nil
}

func (ur *UserRepository) Save(data *User) error {

	if err := ur.db.Debug().Create(&data).Error; err != nil {
		return err
	}
	return nil
}
