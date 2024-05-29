package repository

import (
	"a21hc3NpZ25tZW50/model"
	"errors"

	"gorm.io/gorm"
)

type UserRepository interface {
	Add(user model.User) error
	CheckAvail(user model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepository {
	return &userRepository{db}
}
func (u *userRepository) Add(user model.User) error {
	if err := u.db.Create(&user).Error; err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (u *userRepository) CheckAvail(user model.User) error {
	var foundUser model.User
	if err := u.db.Where("username = ? AND password = ?", user.Username, user.Password).First(&foundUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return err
	}
	return nil // TODO: replace this
}
