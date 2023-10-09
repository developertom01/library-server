package exceptions

import "github.com/developertom01/library-server/app/graphql/model"

func NewEmailAlreadyExistsError(message string) model.EmailAlreadyExistsError {
	return model.EmailAlreadyExistsError{
		Message: message,
	}

}
