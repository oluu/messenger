package messenger

import "time"

// Message interface
type Message interface {
	GetID() string
	GetClientID() string
	GetRequestID() string
	GetBody() []byte
	GetTimestamp() time.Time
}

type message struct {
	ID        string
	ClientID  string
	RequestID string
	Body      []byte
	Timestamp time.Time
}

// NewMessage func
func NewMessage(clientID, requestID string, body []byte) Message {
	return &message{
		ID:        "123",
		ClientID:  clientID,
		RequestID: requestID,
		Body:      body,
		Timestamp: time.Now(),
	}
}

// GetID returns the message's ID value
func (m *message) GetID() string { return m.ID }

// GetClientID returns the message's clientID value
func (m *message) GetClientID() string { return m.ClientID }

// GetRequestID returns the message's requestID value
func (m *message) GetRequestID() string { return m.RequestID }

// GetBody returns the message's body value
func (m *message) GetBody() []byte { return m.Body }

// GetTimestamp returns the message's timestamp value
func (m *message) GetTimestamp() time.Time { return m.Timestamp }
