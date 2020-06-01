package router

import (
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/websocket"
)

type router struct {
	wsHandlers map[string]func(websocket.Connection, []byte)
}

func (r *router) WebsocketRouter(conn websocket.Connection) {

}

func (r *router) GetIdentifier(ctx context.Context) string {
	return ""
}

var Router *router

func init() {
	Router = &router{}
}
