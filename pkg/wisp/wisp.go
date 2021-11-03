package wisp

import (
	"strings"
	"time"
)

var (
	running bool
	stop    chan struct{}

	handlers []*Handler

	addEvent   chan *Event
	addHandler chan *Handler
	delHandler chan *Handler
)

func broadcast(event *Event, handlers []*Handler) {
	event.Time = time.Now()

	for _, hand := range handlers {
		// skip handlers with nil callbacks to avoid errors
		if hand.Callback == nil {
			continue
		}

		for _, tag := range hand.Tags {
			// find matching handler tags
			if strings.HasPrefix(event.Tag, tag) || tag == "*" {
				if hand.Blocking {
					if hand.Callback(event) {
						return
					}
				} else {
					// note that a handler CANNOT consume the event if it is
					// non-blocking
					go hand.Callback(event)
				}
			}
		}
	}
}

func run() {
	running = true
	defer func() { running = false }()

	for {
		select {
		case <-stop:
			return
		case event := <-addEvent:
			go broadcast(event, handlers)
		case add := <-addHandler:
			handlers = append(handlers, add)
		case del := <-delHandler:
			for i, hand := range handlers {
				if hand == del {
					handlers[i] = handlers[len(handlers)-1]
					handlers = handlers[:len(handlers)-1]
				}
			}
		default:
			continue
		}
	}
}

// Init initializes the event loop. Will not initialize twice.
func Init() {
	if !running {
		handlers = []*Handler{}

		stop = make(chan struct{})
		addEvent = make(chan *Event)
		addHandler = make(chan *Handler)
		delHandler = make(chan *Handler)

		go run()
	}
}

// Stop stops the event loop
func Stop() {
	stop <- struct{}{}
}

// Running returns true if the event loop is running
func Running() bool {
	return running
}

// Broadcast creates and broadcasts an event
func Broadcast(event *Event) {
	addEvent <- event
}

// Emit creates and broadcasts an event
func Emit(tag string, data interface{}) {
	Broadcast(&Event{Tag: tag, Data: data})
}

// AddHandler adds a handler to the event loop
func AddHandler(handler *Handler) {
	addHandler <- handler
}

// DelHandler removes a handler from the event loop
func DelHandler(handler *Handler) {
	delHandler <- handler
}

// Handlers returns a list of all handlers
func Handlers() []*Handler {
	return handlers
}
