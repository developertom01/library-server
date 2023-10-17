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
	UserId uint        `gorm:"column:user_id"`
	Path   []uuid.UUID `gorm:"column:path;type:uuid[]"`

	User     User         `gorm:"foreignKey:UserId"`
	Children []FolderItem `gorm:"foreignKey:ParentId"`
}

func (Folder) TableName() string {
	return "folders"
}

type File struct {
	gorm.Model
	Uuid   uuid.UUID   `gorm:"type:uuid;default:gen_random_uuid();unique;"`
	Name   string      `gorm:"not null;column:name"`
	Url    string      `gorm:"not null;column:name"`
	UserId uint        `gorm:"column:user_id"`
	Path   []uuid.UUID `gorm:"column:path;type:uuid[]"`
	User   User        `gorm:"foreignKey:UserId"`
}

func (File) TableName() string {
	return "files"
}

type FolderItem struct {
	gorm.Model
	ParentId      uint `gorm:"column:parent_id;"`
	FileId        uint `gorm:"column:file_id;"`
	ChildFolderId uint `gorm:"column:child_folder_id;"`

	Parent      *Folder `gorm:"foreignKey:ParentId"`
	ChildFolder *Folder `gorm:"foreignKey:ChildFolderId"`
	File        *File   `gorm:"foreignKey:FileId"`
}

func (FolderItem) TableName() string {
	return "folder_items"
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
