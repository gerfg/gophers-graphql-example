package resolver

import (
	"context"

	"github.com/gerfg/gophers-graphql/model"
	"github.com/gerfg/gophers-graphql/repository"
)

func (r *Resolver) GetUser(ctx context.Context, args struct{ Name string }) *model.UserResolver {
	jose := repository.GetUser("josé")
	resolver := model.UserResolver{
		Usr: jose,
	}
	return &resolver
}
