package resources

import (
	"github.com/developertom01/library-server/app/graphql/model"
	"github.com/developertom01/library-server/internals/entities"
	"github.com/developertom01/library-server/utils"
)

func NewUserResource(user entities.User) model.User {
	return model.User{
		FirstName: &user.FirstName,
		LastName:  &user.LastName,
		Email:     user.Email,
		CreatedAt: utils.ConvertTimeToIso(user.CreatedAt),
		UpdatedAt: utils.ConvertTimeToIso(user.UpdatedAt),
	}
}
