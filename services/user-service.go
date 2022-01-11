package services

type UserService interface {
	GetById()
	GetAll()
}

type userService struct {}

func NewUserService() UserService {
	return &userService{}
}

func (u *userService) GetById() {

}

func (u *userService) GetAll() {
	
}
