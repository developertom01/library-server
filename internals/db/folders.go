package db

import (
	"github.com/developertom01/library-server/internals/entities"
	"github.com/google/uuid"
)

func mapFolderItemsToChildFolders(items []entities.FolderItem) []entities.Folder {
	var children []entities.Folder
	for _, item := range items {
		children = append(children, item.ChildFolder)
	}

	return children
}

func (db *Database) FindUsersTopLevelFolders(userId int) ([]entities.Folder, error) {
	rootFolder := entities.Folder{
		UserId: userId,
		IsRoot: true,
	}
	res := db.DB.Find(rootFolder)

	return mapFolderItemsToChildFolders(rootFolder.Children), res.Error
}

func (db *Database) FindChildrenReferencingParentId(parentUuid uuid.UUID) ([]entities.Folder, error) {
	rootFolder := entities.Folder{
		Uuid: parentUuid,
	}
	res := db.DB.Model(&entities.Folder{}).Joins("Children").Joins("Children.ChildFolder").Find(rootFolder)
	return mapFolderItemsToChildFolders(rootFolder.Children), res.Error

}
