package db

import (
	"fmt"
	"sync"

	"github.com/developertom01/library-server/app/graphql/model"
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

func (db *Database) FindUsersTopLevelFolderItems(userId int, limit int, offset int, orderByField *string, orderBy *model.Order) ([]entities.FolderItem, int, error) {
	var totalCount int64
	var err error
	var errMutex sync.Mutex
	var wg sync.WaitGroup

	if orderByField == nil {
		defaultOrderByField := "id"
		orderByField = &defaultOrderByField
	}

	rootFolder := entities.Folder{
		UserId: userId,
		IsRoot: true,
	}
	wg.Add(2)

	go func() {
		defer wg.Done()
		subQuery := db.DB.Model(&model.FolderItem{}).Limit(limit).Offset(offset).Order(fmt.Sprintf("%s %s", *orderByField, orderBy.String()))
		res := db.DB.Model(&model.Folder{}).Joins("Children", subQuery).Find(&rootFolder)
		errMutex.Lock()
		err = res.Error
		errMutex.Unlock()
	}()

	go func() {
		defer wg.Done()
		res := db.DB.Model(&entities.FolderItem{}).
			Joins(`LEFT JOIN folders AS child_folders 
			   ON folder_items.child_folder_id = child_folder.id 
		   WHERE child_folder.user_id=? AND is_root=TRUE`,
				userId).Count(&totalCount)
		errMutex.Lock()
		err = res.Error
		errMutex.Unlock()
	}()

	return rootFolder.Children, int(totalCount), err
}

func (db *Database) FindChildrenReferencingParentId(parentUuid uuid.UUID) ([]entities.FolderItem, error) {
	rootFolder := entities.Folder{
		Uuid: parentUuid,
	}
	subQuery := db.DB.Model(&entities.FolderItem{})
	res := db.DB.Model(&entities.Folder{}).Joins("Children", subQuery).First(rootFolder)

	return rootFolder.Children, res.Error
}

func (db *Database) FindPaginatedFolderContentReferencingParentId(parentUuid uuid.UUID, limit uint, offset uint, orderByField *string, orderBy *model.Order) ([]entities.FolderItem, int, error) {
	var totalCount int64
	var err error
	var wg sync.WaitGroup
	var errMutex sync.Mutex

	if orderByField == nil {
		defaultOrderByField := "id"
		orderByField = &defaultOrderByField
	}

	rootFolder := entities.Folder{
		Uuid: parentUuid,
	}
	wg.Add(2)
	go func() {
		defer wg.Done()
		res := db.DB.Model(&entities.FolderItem{}).
			Joins(`LEFT JOIN folders AS child_folders 
				   	ON folder_items.child_folder_id = child_folder.id 
			       WHERE child_folder.uuid=?`,
				parentUuid).Count(&totalCount)

		errMutex.Lock()
		err = res.Error
		errMutex.Unlock()

	}()
	go func() {
		defer wg.Done()
		subQuery := db.DB.Model(&entities.FolderItem{}).Limit(int(limit)).Offset(int(offset)).Order(fmt.Sprintf("%s %s", *orderByField, orderBy.String()))
		res := db.DB.Model(&entities.Folder{}).Joins("Children", subQuery).First(rootFolder)
		errMutex.Lock()
		err = res.Error
		errMutex.Unlock()
	}()
	wg.Wait()
	return rootFolder.Children, int(totalCount), err
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
