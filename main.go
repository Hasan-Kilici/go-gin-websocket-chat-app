package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
)

type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
	Color string `json:"color"`
}
var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)
var upgrader = websocket.Upgrader{}
var sender *websocket.Conn


func handleConnections(w http.ResponseWriter, r *http.Request) {
	  ws, err := upgrader.Upgrade(w, r, nil)
	  if err != nil {
		log.Fatal(err)
	  }
	  defer ws.Close()
	  clients[ws] = true
	  for {
		var msg Message
		sender = ws
		err := ws.ReadJSON(&msg)
		if err != nil {
		  log.Printf("error: %v", err)
		  delete(clients, ws)
		  break
		}
		broadcast <- msg
	  }
  }
  
  func handleMessages() {
	  for {
		  msg := <-broadcast
		  for client := range clients {
			if(client != sender){
					err := client.WriteJSON(msg)
					if err != nil {
						log.Printf("error: %v", err)
						client.Close()
						delete(clients, client)
					}
			}
		  }
	  }
  }

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("src/*.tmpl")
	r.Static("/static", "./static/")
  
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Web Socket, Chat app",
		})
	})

	r.GET("/ws", func(c *gin.Context) {
		handleConnections(c.Writer, c.Request)
	})
	return r
}

func main() {
	r := setupRouter()
	go handleMessages()
	r.Run(":6000")
}