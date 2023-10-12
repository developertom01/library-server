package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Folder struct {
	gorm.Model
	Uuid   uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();unique;"`
	Name   string    `gorm:"not null;column:name"`
	IsRoot bool      `gorm:"column:is_root;default:false"`
	UserId int       `gorm:"column:user_id"`

	User     User         `gorm:"references:UserId"`
	Children []FolderItem `gorm:"foreignKey:ParentId"`
}

type File struct {
	gorm.Model
	Uuid   uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();unique;"`
	Name   string    `gorm:"not null;column:name"`
	Url    string    `gorm:"not null;column:name"`
	UserId int       `gorm:"column:user_id"`
	User   User      `gorm:"references:UserId"`
}

type FolderItem struct {
	gorm.Model
	ParentId      int `gorm:"column:parent_id;"`
	FileId        int `gorm:"column:file_id;"`
	ChildFolderId int `gorm:"column:child_folder_id;"`

	Parent      Folder `gorm:"references:ParentId"`
	ChildFolder Folder `gorm:"references:ChildFolderId"`
	File        File   `gorm:"references:fileId"`
}
