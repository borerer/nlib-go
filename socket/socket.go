package socket

import (
	"errors"
	"sync"
	"time"

	"github.com/borerer/nlib-go/utils"
	nlibshared "github.com/borerer/nlib-shared/go"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

const (
	defaultWebSocketMessageTimeout = time.Second * 15
)

type Socket struct {
	url                     string
	timeout                 time.Duration
	conn                    *websocket.Conn
	lock                    sync.Mutex
	pendingResponseChannels sync.Map
	requestHandler          func(req *nlibshared.WebSocketMessage) (*nlibshared.WebSocketMessage, error)
	messageChannel          chan *nlibshared.WebSocketMessage
	errorChannel            chan error
}

var (
	ErrTimeout        = errors.New("timeout")
	ErrClosed         = errors.New("closed")
	ErrNoPaired       = errors.New("no paired")
	ErrInvalidMessage = errors.New("invalid message")
)

func NewSocket(url string) *Socket {
	s := &Socket{
		url:     url,
		timeout: defaultWebSocketMessageTimeout,
	}
	s.lock.Lock()
	return s
}

func (s *Socket) SetTimeout(timeout time.Duration) {
	s.timeout = timeout
}

func (s *Socket) SetRequestHandler(handler func(req *nlibshared.WebSocketMessage) (*nlibshared.WebSocketMessage, error)) {
	s.requestHandler = handler
}

func (s *Socket) sendMessage(message *nlibshared.WebSocketMessage) error {
	if err := s.conn.WriteJSON(message); err != nil {
		return err
	}
	return nil
}

func (s *Socket) SendRequest(req *nlibshared.WebSocketMessage) (*nlibshared.WebSocketMessage, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if len(req.MessageID) == 0 {
		req.MessageID = uuid.NewString()
	}
	if len(req.Type) == 0 {
		req.Type = nlibshared.WebSocketMessageTypeRequest
	}
	if req.Timestamp == 0 {
		req.Timestamp = time.Now().UnixMilli()
	}
	ch := make(chan *nlibshared.WebSocketMessage, 1)
	s.pendingResponseChannels.Store(req.MessageID, ch)
	if err := s.sendMessage(req); err != nil {
		return nil, err
	}
	select {
	case res := <-ch:
		return res, nil
	case <-time.After(s.timeout):
		return nil, ErrTimeout
	}
}

func (s *Socket) onMessage(message *nlibshared.WebSocketMessage) {
	switch message.Type {
	case nlibshared.WebSocketMessageTypeRequest:
		if s.requestHandler != nil {
			res, err := s.requestHandler(message)
			if err != nil {
				utils.LogError(err)
				return
			}
			if len(res.MessageID) == 0 {
				res.MessageID = uuid.NewString()
			}
			if len(res.PairMessageID) == 0 {
				res.PairMessageID = message.MessageID
			}
			if len(res.Type) == 0 {
				res.Type = nlibshared.WebSocketMessageTypeResponse
			}
			if res.Timestamp == 0 {
				res.Timestamp = time.Now().UnixMilli()
			}
			err = s.sendMessage(res)
			if err != nil {
				utils.LogError(err)
			}
		}
	case nlibshared.WebSocketMessageTypeResponse:
		if chRaw, ok := s.pendingResponseChannels.LoadAndDelete(message.PairMessageID); ok {
			if ch, ok := chRaw.(chan *nlibshared.WebSocketMessage); ok {
				ch <- message
			} else {
				utils.LogError(errors.New("unexpected channel type"))
			}
		} else {
			utils.LogError(ErrNoPaired)
			utils.PrintJSON(message)
		}
	default:
		utils.LogError(ErrInvalidMessage)
		utils.PrintJSON(message)
	}
}

func (s *Socket) Connect() error {
	var err error
	s.conn, _, err = websocket.DefaultDialer.Dial(s.url, nil)
	if err != nil {
		return err
	}
	s.messageChannel = make(chan *nlibshared.WebSocketMessage)
	s.errorChannel = make(chan error)
	s.conn.SetCloseHandler(func(code int, text string) error {
		s.errorChannel <- ErrClosed
		s.lock.Lock()
		return nil
	})
	s.lock.Unlock()
	return nil
}

func (s *Socket) ListenWebSocketMessages() error {
	go func() {
		for {
			var message nlibshared.WebSocketMessage
			if err := s.conn.ReadJSON(&message); err != nil {
				s.errorChannel <- err
			}
			s.onMessage(&message)
		}
	}()
	return <-s.errorChannel
}
