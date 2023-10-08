package socket

import (
	"log"
	"net/http"
	"strings"

	"github.com/developertom01/library-server/internals/db"
	"github.com/developertom01/library-server/utils"
	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/polling"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
)

type Socket struct {
	Db     *db.Database
	Server *socketio.Server
}

var allowOriginFunc = func(r *http.Request) bool {
	return true
}

func handleError(c socketio.Conn, err error) error {
	if err != nil {
		c.Emit(CONNECTION_ERROR)
		return err
	}
	return nil
}

func NewSocket(db *db.Database) *Socket {
	server := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{
			&polling.Transport{
				CheckOrigin: allowOriginFunc,
			},
			&websocket.Transport{
				CheckOrigin: allowOriginFunc,
			},
		},
	})
	server.OnConnect("/", func(s socketio.Conn) error {
		token := s.RemoteHeader().Get("Authorization")
		_, err := utils.ValidateToken(strings.Replace(token, "Bearer ", "", 1))
		handleError(s, err)
		return nil
	})
	go func() {
		if err := server.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()
	// defer server.Close()
	return &Socket{
		Server: server,
		Db:     db,
	}
}

func (s *Socket) Close() {
	s.Close()
}
