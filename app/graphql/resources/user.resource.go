package resources

import (
	"github.com/developertom01/library-server/app/graphql/model"
	"github.com/developertom01/library-server/internals/entities"
)

func NewUserResource(user entities.User) model.User {
	return model.User{
		FirstName: &user.FirstName,
		LastName:  &user.LastName,
		Email:     user.Email,
		CreatedAt: &user.CreatedAt,
		UpdatedAt: &user.UpdatedAt,
	}
}
