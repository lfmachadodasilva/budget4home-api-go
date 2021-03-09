package user

import (
	"github.com/gin-gonic/gin"
)

// IUserService user service interface
type IUserService interface {
	Fetch(c *gin.Context) (res []User, err error)
	FetchByIds(c *gin.Context, ids *map[string]*User) (err error)
}

// Service user service
type Service struct {
	repo IUserRepository
}

// NewService constructor
func NewService(r IUserRepository) IUserService {
	return &Service{
		repo: r,
	}
}

// Fetch get all users
func (s *Service) Fetch(c *gin.Context) (res []User, err error) {
	res, err = s.repo.Fetch(c)
	return res, err
}

// FetchByIds get all users by IDs
func (s *Service) FetchByIds(c *gin.Context, ids *map[string]*User) (err error) {
	err = s.repo.FetchByID(c, ids)
	return err
}
