package repositories

import (
	"github.com/titoyudha/go_ecommerce_api/model"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	SignUp(user model.User) model.User
	FindByID(id string) model.User
	FindAll() []model.User
	Delete(user model.User)
	// SignIn(password string) string
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) SignUp(user model.User) model.User {
	u.db.Save(&user)
	return user
}

func (u *userRepository) FindByID(id string) model.User {
	var user model.User
	u.db.First(&user, id)
	return user
}

func (u *userRepository) FindAll() []model.User {
	var user []model.User
	u.db.Find(&user)

	return user
}

func (u *userRepository) Delete(user model.User) {
	u.db.Delete(&user)
}

// func (u *userRepository) SignIn(password string) string {

// }
