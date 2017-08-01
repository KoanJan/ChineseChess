package core

/*
数据传输结构(服务端内部)
*/
type Msg struct {
	UID      string // 客户端唯一ID
	Callback string // 回调名称
	Data     []byte // 数据
}

/*
数据传输结构(来自客户端)
*/
type ConnMsg struct {
	Callback string // 回调名称
	Data     []byte // 数据
}
