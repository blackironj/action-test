package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "I'm test")
	})

	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello!!!")
	})

	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "i am okay")
	})

	r.GET("/ws", func(c *gin.Context) {
		wshandler(c.Writer, c.Request)
	})

	r.Run(":8080")
}

func wshandler(w http.ResponseWriter, r *http.Request) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	defer conn.Close()
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: ", err)
		return
	}

	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("disconnected")
			break
		}
		conn.WriteMessage(t, msg)
	}
}
