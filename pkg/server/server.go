package server

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
	"github.com/leprosus/wow/internal/quotes"
)

const resourceLength = 16

type Server struct {
	config *config.Config
	quotes *quotes.Collection
}

func init() {
	gob.Register(hashcash.Hash{})
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
		quotes: quotes.NewCollection(),
	}
}

func (s *Server) ListenAndServe() (err error) {
	addr := net.JoinHostPort(s.config.ServerHost, s.config.ServerPort)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer func() {
		e := listener.Close()
		if err != nil {
			err = fmt.Errorf("%w: %v", err, e)
		} else {
			err = e
		}
	}()

	log.Println("listening", listener.Addr())

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}

		s.handle(conn)
	}
}

func (s *Server) handle(conn net.Conn) {
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	log.Println("new client:", conn.RemoteAddr())

	reader := bufio.NewReader(conn)

	for {
		msg, err := s.readMessage(reader)
		if err != nil {
			log.Println(err)

			continue
		}

		switch msg.Header {
		case protocol.AskType:
			err = s.handleAksRequest(conn)
			if err != nil {
				log.Println(err)
			}

			continue
		case protocol.AnswerType:
			answer, ok := msg.Body.(protocol.Answer)
			if !ok {
				log.Println("receive unexpected a message header")

				return
			}

			err = s.handleAnswerRequest(conn, answer)
			if err != nil {
				log.Println(err)
			}

			return
		default:
			log.Println("receive unsupported a message header")
		}
	}
}

func (s *Server) readMessage(reader *bufio.Reader) (protocol.Message, error) {
	var msg protocol.Message
	err := gob.NewDecoder(reader).Decode(&msg)

	return msg, err
}

func (s *Server) handleAksRequest(conn net.Conn) error {
	resource, err := hashcash.GenRandString(resourceLength)
	if err != nil {
		return err
	}

	var challenge *protocol.Challenge
	challenge, err = hashcash.NewHash(resource, s.config.HashcashZerosAmount)
	if err != nil {
		return err
	}

	msg := protocol.Message{
		Header: protocol.ChallengeType,
		Body:   *challenge,
	}

	bs, err := msg.Encode()
	if err != nil {
		return err
	}

	_, err = conn.Write(bs)

	return err
}

func (s *Server) handleAnswerRequest(conn net.Conn, answer protocol.Answer) error {
	ttl := time.Duration(s.config.ChallengeTTL) * time.Second
	msg := protocol.Message{}

	if answer.Verify() && answer.Date.Add(ttl).After(time.Now()) {
		msg.Header = protocol.GrandType
		msg.Body = s.quotes.GetRandQuote()
	} else {
		msg.Header = protocol.ErrorType
	}

	bs, err := msg.Encode()
	if err != nil {
		return err
	}

	_, err = conn.Write(bs)

	return err
}
