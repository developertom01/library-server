package db

import (
	"fmt"

	"github.com/developertom01/library-server/internals/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

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

func (db *Database) FindUserById(id int) (*entities.User, error) {
	var user entities.User
	res := db.DB.Where("id=?", id).First(&user)

	return &user, res.Error
}

func (db *Database) CreateUser(firstName string, lastName string, email string, password string) (*entities.User, error) {
	user := entities.User{FirstName: firstName, LastName: lastName, Email: email, Password: password}
	res := db.DB.Create(&user)

	return &user, res.Error
}

func (db *Database) CreateUserWithFileSetup(firstName string, lastName string, email string, password string) (*entities.User, error) {
	var user entities.User
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		user = entities.User{FirstName: firstName, LastName: lastName, Email: email, Password: password}

		if res := tx.Create(&user); res.Error != nil {
			return res.Error
		}

		rootFolder := entities.Folder{
			UserId: user.ID,
			Name:   fmt.Sprintf("root_folder_%s", user.Uuid.String()),
			IsRoot: true,
			Path:   []uuid.UUID{},
		}
		if res := tx.Create(&rootFolder); res.Error != nil {
			return res.Error
		}

		return nil

	})

	return &user, err
}
