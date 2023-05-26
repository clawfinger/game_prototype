package event

type subscription struct {
	observers []Observer
}

func (s *subscription) subscribe(obs Observer) {
	s.observers = append(s.observers, obs)
}

func (s *subscription) broadcast(e Event) {
	for _, obs := range s.observers {
		obs.Notify(e)
	}
}
