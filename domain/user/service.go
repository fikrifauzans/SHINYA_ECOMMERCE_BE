package user

type UserService interface {
	GetUsers(page int, perPage int) ([]User, int64, error)
	GetUserByID(id uint) (User, error)
	CreateUser(user User) (User, error)
	UpdateUser(user User) (User, error)
	DeleteUser(id uint) error
}
