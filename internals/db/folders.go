package db

import (
	"github.com/developertom01/library-server/internals/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func mapFolderItemsToChildFolders(items []entities.FolderItem) []entities.Folder {
	var children []entities.Folder
	for _, item := range items {
		children = append(children, *item.ChildFolder)
	}

	return children
}

func (db *Database) FindUsersTopLevelFolderItems(userId int) ([]entities.FolderItem, error) {
	rootFolder := entities.Folder{
		UserId: userId,
		IsRoot: true,
	}
	res := db.DB.Find(rootFolder)

	return rootFolder.Children, res.Error
}

func (db *Database) FindChildrenReferencingParentId(parentUuid uuid.UUID) ([]entities.FolderItem, error) {
	rootFolder := entities.Folder{
		Uuid: parentUuid,
	}
	subQuery := db.DB.Model(&entities.FolderItem{})
	res := db.DB.Model(&entities.Folder{}).Joins("Children", subQuery).First(rootFolder)

	return rootFolder.Children, res.Error
}

func (db *Database) FindPaginatedFolderContentReferencingParentId(parentUuid uuid.UUID, limit uint, offset uint) ([]entities.FolderItem, error) {
	rootFolder := entities.Folder{
		Uuid: parentUuid,
	}
	subQuery := db.DB.Model(&entities.FolderItem{}).Limit(int(limit)).Offset(int(offset))
	res := db.DB.Model(&entities.Folder{}).Joins("Children", subQuery).First(rootFolder)

	return rootFolder.Children, res.Error
}

func (db *Database) GetUserRootFolder(userId int) (entities.Folder, error) {
	folder := entities.Folder{
		UserId: userId,
		IsRoot: true,
	}
	res := db.DB.First(&folder)

	return folder, res.Error
}

func (db *Database) CreateFolder(folderName string, ownerId int, parentId int) (entities.Folder, error) {
	var parentItem entities.FolderItem
	folder := entities.Folder{
		UserId: ownerId,
		Name:   folderName,
	}
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		if res := tx.Create(&folder); res.Error != nil {
			return res.Error
		}
		parentItem = entities.FolderItem{
			ParentId:      parentId,
			ChildFolderId: int(folder.ID),
		}
		res := tx.Create(&parentItem)
		return res.Error
	})

	return folder, err
}

func (db *Database) FindParentFolderReferencingFolderItemId(folderItemId uint) (entities.Folder, error) {
	item := entities.FolderItem{
		Model: gorm.Model{
			ID: folderItemId,
		},
	}
	res := db.DB.Model(&entities.FolderItem{}).Joins("Parent").First(&item)

	return *item.Parent, res.Error
}

func (db *Database) FindFolderReferencingFolderItemId(folderItemId uint) (entities.Folder, error) {
	item := entities.FolderItem{
		Model: gorm.Model{
			ID: folderItemId,
		},
	}
	res := db.DB.Model(&entities.FolderItem{}).Joins("ChildFolder").First(&item)

	return *item.ChildFolder, res.Error
}

func (db *Database) FindFileReferencingFolderItemId(fileId uint) (entities.File, error) {
	item := entities.FolderItem{
		Model: gorm.Model{
			ID: fileId,
		},
	}
	res := db.DB.Model(&entities.FolderItem{}).Joins("File").First(&item)

	return *item.File, res.Error
}

func (db *Database) CreateFile(fileName string, url string, ownerId int, parentId int) (entities.File, error) {
	var parentItem entities.FolderItem
	file := entities.File{
		UserId: ownerId,
		Name:   fileName,
		Url:    url,
	}
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		if res := tx.Create(&file); res.Error != nil {
			return res.Error
		}
		parentItem = entities.FolderItem{
			ParentId: parentId,
			FileId:   int(file.ID),
		}
		res := tx.Create(&parentItem)
		return res.Error
	})

	return file, err
}
