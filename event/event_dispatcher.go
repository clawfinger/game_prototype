package event

type EventDispatcher struct {
	subscriptions map[string]*subscription
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		subscriptions: make(map[string]*subscription),
	}
}

func (d *EventDispatcher) Subscribe(eventName string, obs Observer) {
	sub, ok := d.subscriptions[eventName]
	if !ok {
		sub = &subscription{}
		d.subscriptions[eventName] = sub
	}
	sub.subscribe(obs)
}

func (d *EventDispatcher) Dispatch(e Event) {
	sub, ok := d.subscriptions[e.Name()]
	if ok {
		sub.broadcast(e)
	}
}
