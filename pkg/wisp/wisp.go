package wisp

import "time"

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
			if event.Tag == tag {
				if hand.Blocking {
					if hand.Callback(event.Data) {
						return
					}
				} else {
					// note that a handler CANNOT consume the event if it is
					// non-blocking
					go hand.Callback(event.Data)
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

func Stop() {
	stop <- struct{}{}
}

func Running() bool {
	return running
}

func Broadcast(event *Event) {
	addEvent <- event
}

func AddHandler(handler *Handler) {
	addHandler <- handler
}

func DelHandler(handler *Handler) {
	delHandler <- handler
}

func Handlers() []*Handler {
	return handlers
}
