package db

import "github.com/developertom01/library-server/internals/entities"

func (db *Database) FindUserByEmail(email string) (*entities.User, error) {
	var user entities.User
	res := db.DB.Where("email=?", email).First(&user)
	return &user, res.Error
}

func (db *Database) FindUserByUuid(uuid string) (*entities.User, error) {
	var user entities.User
	res := db.DB.Where("uuid=?", uuid).First(&user)
	return &user, res.Error
}
