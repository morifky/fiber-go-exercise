package models

import (
	"fiber-go-exercise/utils"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint32 `gorm:"primary_key;auto_increment"`
	Username string `gorm:"uniqueIndex;not null"`
	Email    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
}

type UserRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) FindAll() (*[]User, error) {
	users := []User{}

	if err := u.db.Debug().Model(&User{}).Limit(100).Find(&users).Error; err != nil {
		return &[]User{}, err
	}

	return &users, nil
}

func (u *UserRepository) FindByEmail(email string) (*User, error) {
	user := User{}

	if err := u.db.Debug().Model(&User{}).Where("email = ?", email).Take(&user).Error; err != nil {
		return &User{}, err
	}

	return &user, nil
}

func (u *UserRepository) FindOne(id uint32) (*User, error) {
	user := User{}

	if err := u.db.Debug().Model(&User{}).Where("id = ?", id).Take(&user).Error; err != nil {
		return &User{}, err
	}

	return &user, nil
}

func (u *UserRepository) UpdateOne(id uint32, data *User) (*User, error) {
	var err error

	data.Password, err = utils.HashPassword(data.Password)

	if err != nil {
		zap.S().Warn("Unable to hash password, error: ", err)
	}

	//update data
	db := u.db.Debug().Model(&User{}).Where("id = ?", id).Take(&User{}).UpdateColumns(map[string]interface{}{
		"username": data.Username,
		"email":    data.Email,
		"password": data.Password,
	})

	if db.Error != nil {
		return &User{}, db.Error
	}

	//display updated data
	user, err := u.FindOne(id)
	if err != nil {
		return &User{}, err
	}
	return user, nil

}

func (u *UserRepository) DeleteOne(id uint32) (int64, error) {
	db := u.db.Debug().Model(&User{}).Where("id = ?", id).Take(&User{}).Delete(&User{})
	if db.Error != nil {
		return 0, db.Error
	}

	return db.RowsAffected, nil
}

func (u *UserRepository) Save(data *User) (*User, error) {
	if err := u.db.Debug().Create(&data).Error; err != nil {
		return &User{}, err
	}
	return data, nil
}
