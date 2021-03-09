package group

import (
	"budget4home/src/user"

	"github.com/gin-gonic/gin"
)

// IGroupService group service interface
type IGroupService interface {
	Fetch() (res []Group, err error)
	FetchFull(c *gin.Context) ([]FullResponse, error)
	Add(c *gin.Context) ([]Response, error)
}

// Service group service
type Service struct {
	repo        IGroupRepository
	userService user.IUserService
}

// NewService constructor
func NewService(r IGroupRepository, u user.IUserService) IGroupService {
	return &Service{
		repo:        r,
		userService: u,
	}
}

// Fetch get all groups
func (s *Service) Fetch() (res []Group, err error) {
	res, err = s.repo.Fetch()
	return res, err
}

// FetchFull get all groups
func (s *Service) FetchFull(c *gin.Context) ([]FullResponse, error) {
	// get all groups
	var groups, err = s.repo.Fetch()
	if err != nil {
		return nil, err
	}

	// get all users from all groups
	users := make(map[string]*user.User)
	for _, group := range groups {
		for _, user := range group.Users {
			users[user.UserID] = nil
		}
	}

	// get full data from all users
	err = s.userService.FetchByIds(c, &users)
	if err != nil {
		return nil, err
	}

	var res = []FullResponse{}
	for _, group := range groups {
		var groupUsers = []*user.User{}
		for _, user := range group.Users {
			groupUsers = append(groupUsers, users[user.UserID])
		}
		res = append(res, FullResponse{
			ID:    group.ID,
			Name:  group.Name,
			Users: groupUsers,
		})
	}

	users = nil

	return res, err
}

// Add add new group
func (s *Service) Add(c *gin.Context) ([]Response, error) {
	// TODO validate users

	request := new(AddRequest)
	err := c.Bind(request)
	if err != nil {
		return nil, err
	}

	var group = Group{
		Name: request.Name,
	}

	s.repo.Add(&group)

	var users = []UserGroup{}
	for _, user := range request.Users {
		users = append(users, UserGroup{
			GroupID: group.ID,
			UserID:  user,
		})
	}

	s.repo.AddUserGroup(users)
	users = nil
	request = nil

	return nil, nil
}
