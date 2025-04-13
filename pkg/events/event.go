package events

import "time"

type Event struct {
	Name     string
	Payload  interface{}
	DateTime time.Time
}

func NewEvent() *Event {
	return &Event{
		DateTime: time.Now(),
	}
}

func (e *Event) GetName() string {
	return e.Name
}

func (e *Event) GetPayload() interface{} {
	return e.Payload
}

func (e *Event) SetPayload(payload interface{}) {
	e.Payload = payload
}

func (e *Event) SetName(name string) {
	e.Name = name
}

func (e *Event) GetDateTime() time.Time {
	return e.DateTime
}
