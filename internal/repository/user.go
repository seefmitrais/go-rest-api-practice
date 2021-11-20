package repository

import (
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

type User struct {
	gorm.Model
	Email    string `gorm:"unique_index;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
}

func NewUserRepository(database *gorm.DB) *UserRepository {
	return &UserRepository{db: database}
}

func (ur *UserRepository) Save(user *User) error {
	var err error
	err = ur.db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}
