package dataloader

import "github.com/developertom01/library-server/internals/entities"

func (dl DataLoader) LoadUserById(userId int) (*entities.User, error) {
	return dl.Db.FindUserById(userId)
}
