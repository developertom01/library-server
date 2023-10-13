package dataloader

import (
	"context"

	"github.com/developertom01/library-server/internals/db"
)

type DataLoader struct {
	Db *db.Database
}

func NewDataLoader(db *db.Database) *DataLoader {
	return &DataLoader{
		Db: db,
	}
}

// The function ExtractLoaderFromContext extracts a DataLoader from a context.
func ExtractLoaderFromContext(ctx context.Context) *DataLoader {
	return ctx.Value("DataLoader").(*DataLoader)
}
