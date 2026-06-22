package main

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (app *application) openWSHandler(w http.ResponseWriter, r *http.Request) {
	app.handleWS("open", w, r)
}

func (app *application) sppgWSHandler(w http.ResponseWriter, r *http.Request) {
	id := httprouter.ParamsFromContext(r.Context()).ByName("id")

	app.handleWS("sppg:"+id, w, r)
}

func (app *application) handleWS(room string, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	app.hub.JoinRoom(room, conn)
	defer app.hub.LeaveRoom(room, conn)

	for {
		if _, _, err := conn.ReadMessage(); err != nil {
			break
		}
	}
}
