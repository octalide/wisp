package wisp

type Event struct {
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
