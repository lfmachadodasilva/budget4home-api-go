package user

import "github.com/gin-gonic/gin"

type IUserService interface {
	Fetch(c *gin.Context) (res []User, err error)
}

type UserService struct {
	repo IUserService
}

func NewService(r IUserRepository) IUserService {
	return &UserService{
		repo: r,
	}
}

func (this *UserService) Fetch(c *gin.Context) (res []User, err error) {
	res, err = this.repo.Fetch(c)
	if err != nil {
		return nil, err
	}
	return
}
