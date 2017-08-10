package routers

var routerConf = []handler{

	// hello
	{"hello", hello},
}

//hello
func hello(param []byte) []byte {
	return []byte{'f', 'o', 'f', 'f'}
}
