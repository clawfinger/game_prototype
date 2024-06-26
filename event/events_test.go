package event

import "testing"

type TestObserver struct {
	X int
	Y int
}

func (o *TestObserver) Notify(e Event) {
	switch event := e.(type) {
	case *MousePressedEvent:
		o.X = event.X
		o.Y = event.Y
	}
}

func TestEvent(t *testing.T) {
	tests := []struct {
		event     *MousePressedEvent
		expectedX int
		expectedY int
	}{
		{
			event: &MousePressedEvent{
				X: 10,
				Y: 20,
			},
			expectedX: 10,
			expectedY: 20,
		},
	}
	ed := NewEventDispatcher()
	to := &TestObserver{}
	ed.Subscribe(MousePressedEventName, to)
	for _, test := range tests {
		ed.Dispatch(test.event)
		if to.X != test.expectedX || to.Y != test.expectedY {
			t.Errorf("Failed to dispatch event")
		}
	}
}
