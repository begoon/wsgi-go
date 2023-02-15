package wsgi

import (
	"strings"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

var (
	upgrader = websocket.Upgrader{}
)

func WebSocketHandler(c echo.Context) error {
	cid := c.Param("cid")
	log.Info().Str("cid", cid).Msg("connected")

	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	err = ws.WriteMessage(
		websocket.TextMessage,
		[]byte("WebSocket Server "+Version),
	)
	if err != nil {
		log.Err(err).Msg("write message failed")
		return err
	}

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Err(err).Msg("read message failed")
		}
		log.Info().Str("cid", cid).Str("data", string(msg)).Msg("RECEIVED")

		err = ws.WriteMessage(websocket.TextMessage, []byte("ACK"))
		if err != nil {
			log.Err(err).Msg("write ACK message failed")
			return err
		}

		if strings.Contains(string(msg), "BYE") {
			log.Info().Str("cid", cid).Msg("disconnected")
			return nil
		}
	}
}
