package group

import "budget4home/src/user"

// Response group response
type Response struct {
	ID   uint
	Name string
}

// FullResponse group full response
type FullResponse struct {
	ID   uint
	Name string

	Users []*user.User
}
