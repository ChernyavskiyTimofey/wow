package client

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/leprosus/wow/internal/config"
	"github.com/leprosus/wow/internal/hashcash"
	"github.com/leprosus/wow/internal/protocol"
)

var (
	ErrUnexpectedHeader = fmt.Errorf("receive unexpected a message header")
	ErrUnexpectedBody   = fmt.Errorf("receive unexpected a message body")
)

type Client struct {
	config *config.Config
}

func init() {
	gob.Register(hashcash.Hash{})
}

func NewClient(config *config.Config) *Client {
	return &Client{
		config: config,
	}
}

func (c *Client) Run() {
	for range time.NewTicker(3 * time.Second).C {
		err := c.run()
		if err != nil {
			log.Println(err)
		}
	}
}

func (c *Client) run() error {
	addr := net.JoinHostPort(c.config.ServerHost, c.config.ServerPort)

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}

	log.Println("connected to", addr)

	defer func() {
		e := conn.Close()
		if err != nil {
			err = fmt.Errorf("%w: %v", err, e)
		} else {
			err = e
		}
	}()

	msg, err := c.doRequest(conn)
	if err != nil {
		log.Println(err)

		return err
	}

	fmt.Println("quote:", msg)

	return nil
}

func (c *Client) doRequest(conn net.Conn) (string, error) {
	challenge, err := c.doAskRequest(conn)
	if err != nil {
		return "", err
	}

	return c.doAnswerRequest(conn, challenge)
}

func (c *Client) doAskRequest(conn net.Conn) (protocol.Challenge, error) {
	msg := protocol.Message{
		Header: protocol.AskType,
		Body:   nil,
	}

	bs, err := msg.Encode()
	if err != nil {
		return protocol.Challenge{}, err
	}

	_, err = conn.Write(bs)
	if err != nil {
		return protocol.Challenge{}, err
	}

	reader := bufio.NewReader(conn)

	msg, err = c.readMessage(reader)
	if err != nil {
		return protocol.Challenge{}, err
	}

	if msg.Header != protocol.ChallengeType {
		return protocol.Challenge{}, ErrUnexpectedHeader
	}

	challenge, ok := msg.Body.(protocol.Challenge)
	if !ok {
		return protocol.Challenge{}, ErrUnexpectedBody
	}

	_, err = challenge.Compute(c.config.HashcashMaxIterations)
	if !ok {
		return protocol.Challenge{}, err
	}

	return challenge, nil
}

func (c *Client) doAnswerRequest(conn net.Conn, challenge protocol.Challenge) (string, error) {
	msg := protocol.Message{
		Header: protocol.AnswerType,
		Body:   challenge,
	}

	bs, err := msg.Encode()
	if err != nil {
		return "", err
	}

	_, err = conn.Write(bs)
	if err != nil {
		return "", err
	}

	reader := bufio.NewReader(conn)

	msg, err = c.readMessage(reader)
	if err != nil {
		return "", err
	}

	if msg.Header != protocol.GrandType {
		return "", ErrUnexpectedHeader
	}

	quote, ok := msg.Body.(string)
	if !ok {
		return "", ErrUnexpectedBody
	}

	return quote, nil
}

func (c *Client) readMessage(reader *bufio.Reader) (protocol.Message, error) {
	var msg protocol.Message
	err := gob.NewDecoder(reader).Decode(&msg)

	return msg, err
}
