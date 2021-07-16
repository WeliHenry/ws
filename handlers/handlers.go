package handlers

import (
	"fmt"
	"github.com/CloudyKit/jet/v6"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var views = jet.NewSet(
	jet.NewOSFileSystemLoader("./templates/html"),
	jet.InDevelopmentMode(), // remove in production
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WsConnection struct {
	Conn *websocket.Conn
}

type WsJsonResponse struct {
	Action      string `json:"action"`
	Message     string `json:"message"`
	MessageType string `json:"message_type"`
	Conn WsConnection `json:"_"`
}



func Home(w http.ResponseWriter, r *http.Request) {
	err := renderPage(w, "home.jet", nil)
	if err != nil {
		log.Println(err)
	}

}

func renderPage(w http.ResponseWriter, tmpl string, data jet.VarMap) error {
	views, err := views.GetTemplate(tmpl)
	if err != nil {
		log.Println(err)
		return err
	}
	views.Execute(w, data, nil)
	return nil

}


func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("connection to websocket was successful")

	var response WsJsonResponse

	response.Message = `<em><small>connected to the server</small></em>`

	err = ws.WriteJSON(response)
	if err != nil {
		log.Println(err)
	}

}
