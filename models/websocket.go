package models

import "github.com/gorilla/websocket"

type WsConnection struct {
	Conn *websocket.Conn
}

type WsPayload struct {
	Action string `json:"action"`
	Username string `json:"username"`
	Message string `json:"message"`
	Conn WsConnection `json:"_"`
}


