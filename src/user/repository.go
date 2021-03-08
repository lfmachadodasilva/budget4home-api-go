package user

import (
	"log"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
)

type IUserRepository interface {
	Fetch(c *gin.Context) (res []User, err error)
}

type UserRepository struct {
	firebaseAuth *auth.Client
}

func NewRepository(firebaseAuth *auth.Client) IUserRepository {
	return &UserRepository{
		firebaseAuth: firebaseAuth,
	}
}

func (this *UserRepository) Fetch(c *gin.Context) ([]User, error) {
	var usersResult []User = []User{}
	var errResult error = nil
	iter := this.firebaseAuth.Users(c, "")
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
