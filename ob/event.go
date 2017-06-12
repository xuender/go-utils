package ob

type Event struct {
	Id     uint32
	Remove bool
	ChOut  chan interface{}
}
