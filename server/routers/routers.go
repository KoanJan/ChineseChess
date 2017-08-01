package routers

type router struct {
	routers map[string]func([]byte) []byte
}

type handler struct {
	name     string
	_handler func([]byte) []byte
}

func (this *router) config(handler ...handler) *router {
	for _, h := range handler {
		this.routers[h.name] = h._handler
	}
	return this
}

func Router() map[string]func([]byte) []byte {

	r := &router{routers: make(map[string]func([]byte) []byte)}

	// config routers
	r.config(

		// helloworld demo
		handler{"hello_world", helloworld},
		handler{"hello_world2", helloworld2},
	)

	return r.routers

}

func helloworld(param []byte) []byte {
	return []byte{'f', 'o', 'f', 'f'}
}

func helloworld2(param []byte) []byte {
	return []byte{'f', 'o', 'f', 'f', '2'}
}
