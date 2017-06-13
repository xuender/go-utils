package ob

type Suck struct {
	Id     uint32
	Remove bool
	ChData chan interface{}
}
