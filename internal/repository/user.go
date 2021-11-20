package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/seefmitrais/go-rest-api-practice/internal/postgresql"
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

func (u *User) Save() error {
	var err error
	err = postgresql.DB.Create(&u).Error
	if err != nil {
		return err
	}
	return nil
}
