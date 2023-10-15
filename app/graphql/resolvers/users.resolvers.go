package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.35

import (
	"context"
	"fmt"

	"github.com/developertom01/library-server/app/graphql/exceptions"
	"github.com/developertom01/library-server/app/graphql/model"
	"github.com/developertom01/library-server/app/graphql/resources"
	"github.com/developertom01/library-server/utils"
)

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.LoginInInput) (model.LoginResponse, error) {
	user, err := r.Db.FindUserByEmail(input.Email)
	if err != nil {
		return nil, fmt.Errorf("User with email does not exist")
	}
	isPasswordConfirmed := utils.ComparePasswords(user.Password, input.Password)
	if !isPasswordConfirmed {
		return nil, fmt.Errorf("Invalid password")
	}
	jwtToken, err := utils.SignToken(utils.JWTClaim{ID: user.ID, UUID: user.Uuid.String(), Email: user.Email, FirstName: user.FirstName, LastName: user.LastName})
	if err != nil {
		return nil, err
	}
	return &model.LoginSuccessResponse{
		AccessToken:  jwtToken.AccessToken,
		RefreshToken: jwtToken.RefreshToken,
	}, nil
}

// SignUp is the resolver for the signUp field.
func (r *mutationResolver) SignUp(ctx context.Context, input *model.SignUpInput) (model.SignUpUserResponse, error) {
	user, err := r.Db.CreateUser(*input.FirstName, *input.LastName, input.Email, input.Password)
	if err != nil && utils.IsUniqueConstraintViolated(err) {
		return exceptions.NewEmailAlreadyExistsError("User with same email exists"), nil
	}
	return resources.NewUserResource(*user), err
}

// CurrentUser is the resolver for the currentUser field.
func (r *queryResolver) CurrentUser(ctx context.Context) (model.CurrentUserResponse, error) {
	claim := ctx.Value("user")
	if claim == nil {
		return exceptions.NewUnAuthorizeError("UnAuthorized"), nil
	}
	jwtClaim := claim.(*utils.JWTClaim)
	user, err := r.Db.FindUserByUuid(jwtClaim.UUID)
	if err != nil {
		return nil, err
	}
	return resources.NewUserResource(*user), nil
}
