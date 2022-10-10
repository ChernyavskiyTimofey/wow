package protocol

import (
	"bytes"
	"encoding/gob"

	"github.com/leprosus/wow/internal/hashcash"
)

type Header uint

const (
	AskType Header = iota
	ChallengeType
	AnswerType
	GrandType
	ErrorType
)

type Message struct {
	Header Header
	Body   any
}

func NewMessage(header Header, body []byte) *Message {
	return &Message{
		Header: header,
		Body:   body,
	}
}

func ParseMessage(bs []byte) (*Message, error) {
	m := &Message{}
	err := gob.NewDecoder(bytes.NewReader(bs)).Decode(m)

	return m, err
}

func (m Message) Encode() ([]byte, error) {
	var buf bytes.Buffer
	err := gob.NewEncoder(&buf).Encode(m)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type Challenge = hashcash.Hash

type Answer = hashcash.Hash

type Grand struct {
	Quite string
}
