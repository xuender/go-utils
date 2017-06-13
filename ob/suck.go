package ob

type Suck struct {
	Id     uint32
	Close  bool
	ChData chan interface{}
}
