package user

import (
	"log"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
)

// IUserRepository user repository interface
type IUserRepository interface {
	Fetch(c *gin.Context) (res []User, err error)
	FetchByID(c *gin.Context, ids *map[string]*User) (err error)
}

// Repository user repository
type Repository struct {
	firebaseAuth *auth.Client
}

// NewRepository constructor
func NewRepository(firebaseAuth *auth.Client) IUserRepository {
	return &Repository{
		firebaseAuth: firebaseAuth,
	}
}

// Fetch return all users
func (r *Repository) Fetch(c *gin.Context) ([]User, error) {
	var usersResult []User = []User{}
	var errResult error = nil
	iter := r.firebaseAuth.Users(c, "")
	for {
		user, err := iter.Next()
		if err == iterator.Done {
			errResult = err
			break
		}
		if err != nil {
			log.Fatalf("error listing users: %s\n", err)
		}
		usersResult = append(usersResult, User{
			Id:       user.UID,
			Name:     user.DisplayName,
			Email:    user.Email,
			PhotoUrl: user.PhotoURL,
		})
	}

	return usersResult, errResult
}

// FetchByID return all users by ID
func (r *Repository) FetchByID(c *gin.Context, ids *map[string]*User) error {
	var tmp = []auth.UserIdentifier{}

	for id := range *ids {
		tmp = append(tmp, auth.UIDIdentifier{UID: id})
	}

	getUsersResult, err := r.firebaseAuth.GetUsers(c, tmp)
	if err != nil {
		return err
	}

	for _, user := range getUsersResult.Users {
		(*ids)[user.UID] = &User{
			Id:       user.UID,
			Name:     user.DisplayName,
			Email:    user.Email,
			PhotoUrl: user.PhotoURL,
		}
	}

	return nil
}
