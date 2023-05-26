package event

const (
	MousePressedEventID = "MousePressedEvent"
)

type MousePressedEvent struct {
	x int
	y int
}

func (e *MousePressedEvent) Name() string {
	return "MousePressedEvent"
}
