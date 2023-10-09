package exceptions

import "github.com/developertom01/library-server/app/graphql/model"

func NewUnAuthorizeError(message string) model.UnAuthorizedError {
	return model.UnAuthorizedError{
		Message: message,
	}

}
