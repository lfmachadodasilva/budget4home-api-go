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
	// firebaseAuth := c.MustGet("FirebaseAuth").(*auth.Client)

	iter := this.firebaseAuth.Users(c, "")
	for {
		user, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("error listing users: %s\n", err)
		}
		log.Printf("read user user: %v\n", user)
	}

	return nil, nil
}
