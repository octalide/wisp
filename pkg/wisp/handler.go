package wisp

// A Callback will be executed when an event's tags match one of a handlers.
// If the callback returns true, the event will be consumed.
// ONLY blocking handlers can consume events.
type Callback func(e *Event) bool

type Handler struct {
	Callback
	Tags []string

	// A Blocking handler will not have its callback executed in a new goroutine
	// and will be capable of consuming the event.
	Blocking bool
}
