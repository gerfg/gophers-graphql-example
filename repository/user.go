package repository

import "github.com/gerfg/gophers-graphql/model"

func AllUsers() []model.User {
	users := make([]model.User, 0)

	users = append(users, model.NewJoseJunior())
	users = append(users, model.NewJoseJunior())
	users = append(users, model.NewJoseJunior())

	return users
}

func GetUser(firstName string) model.User {
	// DB operation to get the uer with that firstname
	return model.NewJoseJunior()
}
