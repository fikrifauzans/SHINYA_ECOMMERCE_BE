package user

import (
	"app/src/domain/user"
	"app/src/infrastructure/config"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository() user.UserRepository {
	return &userRepository{
		db: config.DB,
	}
}

func (r *userRepository) FindAll(offset int, limit int) ([]user.User, int64, error) {
	var users []user.User
	var total int64
	r.db.Offset(offset).Limit(limit).Find(&users)
	r.db.Model(&user.User{}).Count(&total)
	return users, total, nil
}

func (r *userRepository) FindByID(id uint) (user.User, error) {
	var user user.User
	if err := r.db.First(&user, id).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) Create(user user.User) (user.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) Update(user user.User) (user.User, error) {
	if err := r.db.Save(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) Delete(id uint) error {
	if err := r.db.Delete(&user.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
