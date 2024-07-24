package user

type UserRepository interface {
	FindAll(offset int, limit int) ([]User, int64, error)
	FindByID(id uint) (User, error)
	Create(user User) (User, error)
	Update(user User) (User, error)
	Delete(id uint) error
	FindByEmail(email string) (User, error)
}
