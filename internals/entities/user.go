package entities

import (
	"github.com/developertom01/library-server/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Uuid      uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	FirstName string    `gorm:"column:first_name"`
	LastName  string    `gorm:"column:last_name"`
	Email     string    `gorm:"not null; unique;"`
	Password  string    `gorm:"not null;"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password, err = utils.HashPassword(u.Password)
	return err
}
