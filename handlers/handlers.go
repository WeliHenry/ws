package handlers

import (
	"fmt"
	"github.com/CloudyKit/jet/v6"
	"github.com/gorilla/websocket"
	"github.com/welihenry/ws/models"
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







var wsChan = make(chan models.WsPayload)
var clients = make(map[models.WsConnection]string)


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

	var response models.WsJsonResponse

	response.Message = `<em><small>connected to the server</small></em>`

	conn:= models.WsConnection{
		Conn: ws,
	}
	clients[conn]= ""

	err = ws.WriteJSON(response)
	if err != nil {
		log.Println(err)
	}
	go ListenForWs(&conn)

}


func ListenForWs(conn *models.WsConnection)  {
	defer func() {
		if r := recover(); r != nil{
			log.Println("Error" , fmt.Sprintf("%v", r))
		}
	}()
	var payload models.WsPayload

	for  {
		err:= conn.Conn.ReadJSON(&payload)
		if err != nil {
			//
		}else {
			payload.Conn = *conn
			wsChan <- payload
		}
	}

}

func ListenToWsChan()  {
	var response models.WsJsonResponse
	for  {
		e:= <- wsChan
		response.Action = "got here"
		response.Message = fmt.Sprintf("this is some message, and action was  %s", e.Action)
		brodcastWs(response)
	}
}

func brodcastWs (resp models.WsJsonResponse)  {
	for client := range clients {
		err:= client.Conn.WriteJSON(&resp)
		if err != nil {
			log.Println("websocket error")
			_ = client.Conn.Close()
			delete(clients, client)
		}
	}
}




