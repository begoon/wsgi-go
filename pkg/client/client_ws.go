package client

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

func WebSocketClient() {
	cid := os.Args[2]
	host := os.Args[3]
	url := strings.TrimSuffix(host, "/") + "/api/ws/" + cid
	fmt.Println("- connecting to", url)

	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	t, message, err := c.ReadMessage()
	if err != nil {
		panic(err)
	}
	fmt.Println("RECEIVED", t, string(message))

	q := []string{"I'm " + cid, "Frame 1", "Frame 2", "Frame 3", "BYE"}
	for _, msg := range q {
		if err := c.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
			panic(err)
		}
		fmt.Println("SENT", string(msg))
		t, ack, err := c.ReadMessage()
		if err != nil {
			panic(err)
		}
		fmt.Println("RECEIVED", t, string(ack))
		time.Sleep(time.Duration(1) * time.Second)
	}
}
