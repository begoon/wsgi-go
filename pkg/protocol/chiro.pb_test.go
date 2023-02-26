package protocol

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

func Must[T any](r T, err error) T {
	if err != nil {
		panic(err)
	}
	return r
}

func TestEnvelope(t *testing.T) {
	e := Envelope{
		SequenceNumber: 1,
		EventType:      EventType_CLIENT_START,
		Payload:        &Envelope_ClientStart{&ClientStart{UserID: "abc"}},
	}
	b := Must(proto.Marshal(&e))
	assert.Equal(t, 11, len(b))

	r := Envelope{}
	assert.NoError(t, proto.Unmarshal(b, &r))
	c := r.GetClientStart()
	assert.Equal(t, "abc", c.UserID)
}

func TestClientStart(t *testing.T) {
	m := ClientStart{UserID: "abc"}
	assert.Equal(t, "abc", m.UserID)

	b := Must(proto.Marshal(&m))
	assert.Equal(t, 5, len(b))
	assert.Equal(t, []byte("\x0a\x03abc"), b)

	r := ClientStart{}
	assert.NoError(t, proto.Unmarshal(b, &r))
	assert.Equal(t, "abc", r.UserID)
}
