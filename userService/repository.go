package userService

import "gorm.io/gorm"

type UserRepository interface {
	CreateUser(user User) (User, error)
	GetAllUser() ([]User, error)
	UpdateUser(id uint, updateduser User) (User, error)
	DeleteUser(id uint) error
}
type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}
func (r *userRepository) CreateUser(user User) (User, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}
func (r *userRepository) GetAllUser() ([]User, error) {
	var user []User
	err := r.db.Find(&user).Error
	return user, err
}
func (r *userRepository) UpdateUser(id uint, updateduser User) (User, error) {
	var user User
	result := r.db.First(&user, id)
	if result.Error != nil {
		return User{}, result.Error
	}
	result = r.db.Model(&user).Where("id = ?", id).Updates(updateduser)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}
func (r *userRepository) DeleteUser(id uint) error {
	result := r.db.Delete(&User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
