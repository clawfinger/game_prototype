package event

type Event interface {
	Name() string
}

type Observer interface {
	Notify(e Event)
}
