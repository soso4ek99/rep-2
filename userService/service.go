package userService

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}
func (us *UserService) CreateUser(user User) (User, error) {
	return us.repo.CreateUser(user)
}
func (us *UserService) UpdateUser(id uint, updateduser User) (User, error) {
	return us.repo.UpdateUser(id, updateduser)
}
func (us *UserService) DeleteUser(id uint) error {
	return us.repo.DeleteUser(id)
}
func (us *UserService) GetAllUser() ([]User, error) {
	return us.repo.GetAllUser()
}
