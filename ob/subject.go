package ob

type Subject struct {
	Observers map[uint32]chan interface{}
	ChEvent   chan Event
}

func (s *Subject) Notify(data interface{}) bool {
	if len(s.Observers) == 0 {
		return false
	}
	for _, ch := range s.Observers {
		ch <- data
	}
	return true
}
