package model

import "context"

type User struct {
	FirstName string
	LastName  string
}

func NewUser() User {
	return User{}
}

func NewJoseJunior() User {
	jose := NewUser()

	jose.FirstName = "José"
	jose.LastName = "Junior"

	return jose
}

type UserResolver struct {
	Usr User
}

func NewUserResolver() UserResolver {
	return UserResolver{}
}

func (r *UserResolver) FirstName(ctx context.Context) string {
	return r.Usr.FirstName
}

func (r *UserResolver) LastName(ctx context.Context) string {
	return r.Usr.LastName
}
