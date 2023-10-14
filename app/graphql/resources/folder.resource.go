package resources

import (
	"github.com/developertom01/library-server/app/graphql/model"
	"github.com/developertom01/library-server/internals/entities"
	"github.com/developertom01/library-server/utils"
)

func resolveFolderItemTypeField(fi entities.FolderItem) model.ItemType {
	if fi.ChildFolderId != 0 {
		return model.ItemTypeFile
	}
	return model.ItemTypeFolder
}

func NewFolderResource(folder entities.Folder) *model.Folder {
	return &model.Folder{
		ID:        int(folder.ID),
		UUID:      utils.ParseNativeUuidToScalerUuid(folder.Uuid),
		IsRoot:    folder.IsRoot,
		Name:      folder.Name,
		CreatedAt: utils.ConvertTimeToIso(folder.CreatedAt),
		UpdatedAt: utils.ConvertTimeToIso(folder.UpdatedAt),
		Path:      utils.ParseArrayOfNativeUuidToScalerUuids(folder.Path),
	}
}

func NewFileResource(file entities.File) *model.File {
	return &model.File{
		ID:        int(file.ID),
		UUID:      utils.ParseNativeUuidToScalerUuid(file.Uuid),
		Name:      &file.Name,
		URL:       file.Name,
		CreatedAt: utils.ConvertTimeToIso(file.CreatedAt),
		UpdatedAt: utils.ConvertTimeToIso(file.CreatedAt),
	}
}

func NewFolderResourceCollectionResource(folders []entities.Folder) []*model.Folder {
	var folderCollectionResource []*model.Folder
	for _, folder := range folders {
		folderCollectionResource = append(folderCollectionResource, NewFolderResource(folder))
	}

	return folderCollectionResource
}

func NewFileResourceCollectionResource(files []entities.File) []*model.File {
	var fileCollectionResource []*model.File
	for _, file := range files {
		fileCollectionResource = append(fileCollectionResource, NewFileResource(file))
	}

	return fileCollectionResource
}

func NewFolderItemResource(folderItem entities.FolderItem) *model.FolderItem {
	return &model.FolderItem{
		ID:   int(folderItem.ID),
		Type: resolveFolderItemTypeField(folderItem),
	}
}

func NewFolderItemCollectionResource(contents []entities.FolderItem) []*model.FolderItem {
	var folderContentResource []*model.FolderItem
	for _, item := range contents {
		folderContentResource = append(folderContentResource, NewFolderItemResource(item))
	}

	return folderContentResource
}

func PaginatedFolderItemResource(list []entities.FolderItem, count int, pageSize int, next int, orderByField string, orderBy model.Order) *model.PaginatedFolderItems {
	return &model.PaginatedFolderItems{
		Meta: &model.PaginatedMeta{
			NextPage:     &next,
			PageSize:     pageSize,
			Count:        count,
			OrderBy:      &orderBy,
			OrderByField: &orderByField,
		},
		List: NewFolderItemCollectionResource(list),
	}
}
