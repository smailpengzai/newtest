package ws

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

func GinTest() {
	r := gin.Default()

	r.LoadHTMLFiles("./index.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	r.GET("/ws", func(c *gin.Context) {
		wshandler(c.Writer, c.Request)
	})

	r.Run("localhost:12312")
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wshandler(w http.ResponseWriter, r *http.Request) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(fmt.Sprintf("Failed to set websocket upgrade: %+v", err))
		return
	}

	for {
		t, msg, err := conn.ReadMessage()
		fmt.Println(t, string(msg), err)
		if err != nil {
			break
		}
		if len(msg) != 0 {
			conn.WriteMessage(t, []byte("是的，我活著！"))

		} else {
			conn.WriteMessage(t, []byte("你去哪裡了"))
		}
	}
}
