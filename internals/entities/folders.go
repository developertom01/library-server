package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Folder struct {
	gorm.Model
	Uuid   uuid.UUID   `gorm:"type:uuid;default:gen_random_uuid();unique;"`
	Name   string      `gorm:"not null;column:name"`
	IsRoot bool        `gorm:"column:is_root;default:false"`
	UserId int         `gorm:"column:user_id"`
	Path   []uuid.UUID `gorm:"column:path;default:[]"`

	User     User         `gorm:"references:UserId"`
	Children []FolderItem `gorm:"foreignKey:ParentId"`
}

type File struct {
	gorm.Model
	Uuid   uuid.UUID   `gorm:"type:uuid;default:gen_random_uuid();unique;"`
	Name   string      `gorm:"not null;column:name"`
	Url    string      `gorm:"not null;column:name"`
	UserId int         `gorm:"column:user_id"`
	Path   []uuid.UUID `gorm:"column:path;default:[]"`

	User User `gorm:"references:UserId"`
}

type FolderItem struct {
	gorm.Model
	ParentId      int `gorm:"column:parent_id;"`
	FileId        int `gorm:"column:file_id;"`
	ChildFolderId int `gorm:"column:child_folder_id;"`

	Parent      *Folder `gorm:"references:ParentId"`
	ChildFolder *Folder `gorm:"references:ChildFolderId"`
	File        *File   `gorm:"references:fileId"`
}

func (folder *Folder) BeforeCreate(tx *gorm.DB) (err error) {
	if folder.Uuid == uuid.Nil {
		folder.Uuid = uuid.New()
	}
	return err
}

func (file *File) BeforeCreate(tx *gorm.DB) (err error) {
	if file.Uuid == uuid.Nil {
		file.Uuid = uuid.New()
	}
	return err
}
