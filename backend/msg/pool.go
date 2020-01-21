package msg

type Pool struct {
	Register   chan *ClientMsg
	Unregister chan *ClientMsg
	Clients    map[*ClientMsg]bool
	BroadCast  chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *ClientMsg),
		Unregister: make(chan *ClientMsg),
		Clients:    make(map[*ClientMsg]bool),
		BroadCast:  make(chan Message),
	}
}
