package handler

import (
	"strings"

	pb "github.com/begoon/wsgi-go/pkg/protocol"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/proto"
)

var (
	upgrader = websocket.Upgrader{}
)

func WebSocketTextHandler(c echo.Context) error {
	log.Info().Msg("incoming call")

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

func ProcessMessage(req []byte) ([]byte, error) {
	log.Info().Msg("processing message")
	defer log.Info().Msg("done processing message")

	e := pb.Envelope{}
	if err := proto.Unmarshal(req, &e); err != nil {
		log.Error().Err(err).Msg("failed on unmarshal response")
		return nil, err
	}

	t := e.EventType
	log.Info().Msgf("event type %v", t)

	switch t {
	case pb.EventType_CLIENT_START:
		log.Info().Any("data", e.GetClientStart()).Msg("start")
	case pb.EventType_PAD_FRAMES:
		f := e.GetPadFrame()
		log.Info().Int("image size", len(f.Image)).Msg("frame")
	default:
		log.Error().Msg("unrecognized event type")
	}

	ack := pb.AckMessage{SequenceNumber: e.SequenceNumber + 1000}
	r := pb.Envelope{
		SequenceNumber: 1,
		EventType:      pb.EventType_ACK,
		Payload:        &pb.Envelope_AckMessage{AckMessage: &ack},
	}

	res, err := proto.Marshal(&r)
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal response")
		return nil, err
	}
	return res, nil
}

func WebSocketHandler(c echo.Context) error {
	cid := c.Param("cid")
	log.Info().Str("cid", cid).Msg("connected")

	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Err(err).Msg("unable to upgrade to websocket")
		return err
	}
	defer ws.Close()

	for {
		_, req, err := ws.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
				return nil
			}
			log.Err(err).Msg("read request failed")
			return err
		}
		log.Info().Str("cid", cid).Int("size", len(req)).Msg("RECEIVED")

		resp, err := ProcessMessage(req)
		if err != nil {
			log.Err(err).Msg("request processing failed")
			return err
		}

		log.Info().Str("cid", cid).Int("size", len(resp)).Msg("SENT")

		err = ws.WriteMessage(websocket.BinaryMessage, resp)
		if err != nil {
			log.Err(err).Msg("write response failed")
			return err
		}
	}
}
