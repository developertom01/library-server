package resolvers

import (
	"github.com/developertom01/library-server/app/socket"
	"github.com/developertom01/library-server/internals/db"
	"github.com/developertom01/library-server/internals/object"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Db     *db.Database
	Socket *socket.Socket
	Os     *object.ObjectStorage
}

func NewResolver(db *db.Database, socket *socket.Socket, os *object.ObjectStorage) *Resolver {
	return &Resolver{
		Db:     db,
		Socket: socket,
		Os:     os,
	}
}
