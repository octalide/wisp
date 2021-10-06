package wisp

import "time"

type Event struct {
	time.Time
	Tag  string
	Data interface{}
}

func NewEvent(tag string, data interface{}) *Event {
	e := &Event{
		Tag:  tag,
		Data: data,
	}

	return e
}
