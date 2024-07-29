package usecase

import "app/src/domain/user"

type userUsecase struct {
	userRepository user.UserRepository
}

func NewUserUsecase(r user.UserRepository) user.UserService {
	return &userUsecase{
		userRepository: r,
	}
}

func (u *userUsecase) GetUsers(page int, perPage int) ([]user.User, int64, error) {
	offset := (page - 1) * perPage
	return u.userRepository.FindAll(offset, perPage)
}

func (u *userUsecase) GetUserByID(id uint) (user.User, error) {
	return u.userRepository.FindByID(id)
}

func (u *userUsecase) CreateUser(user user.User) (user.User, error) {
	return u.userRepository.Create(user)
}

func (u *userUsecase) UpdateUser(user user.User) (user.User, error) {
	return u.userRepository.Update(user)
}

func (u *userUsecase) DeleteUser(id uint) error {
	return u.userRepository.Delete(id)
}

func (u *userUsecase) GetUserByEmail(email string) (user.User, error) {
	return u.userRepository.FindByEmail(email)
}
