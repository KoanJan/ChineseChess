package core

import (
	"fmt"
	"net"
	"os"
)

//var ech chan error = make(chan error, 10)    // error
var wch chan []byte = make(chan []byte, 100) // write
var rch chan []byte = make(chan []byte, 100) // read

// tcp服务器
type tcpServer struct {
	ip   string // ip
	port int    // 端口

	conns map[string]*net.TCPConn // 存储连接

	cbs map[string]func([]byte) []byte // 回调集合(由于未加入数据协议, 此处暂时还未有数据解析)
}

// 开启
func (this *tcpServer) start() error {

	listener, err := net.ListenTCP("tcp", &net.TCPAddr{net.ParseIP(this.ip), this.port, ""})
	if err != nil {
		fmt.Printf("服务器错误(详细信息: %v)\n", err)
		return err
	}

	fmt.Printf("TCP服务器成功启动![listen on %d]\n", this.port)

	// 监听wch
	go func() {
		for {
			select {
			case <-wch:
				// TODO 解析数据, 向指定uid推送数据
			}
		}
	}()

	// 监听是否有新的客户端连接
	for {
		func() {
			conn, err := listener.AcceptTCP()
			if err != nil {
				//ech <- err
				fmt.Printf("服务器错误(详细信息: %v)\n", err)
				os.Exit(0)
			}
			conn.SetNoDelay(true)
			uid := make([]byte, 24)
			// 连接第一条信息为24位的唯一id
			if l, err := conn.Read(uid); l != 24 || err != nil {
				//ech <- err
				fmt.Printf("连接错误(uid=%s 详细信息: %v)\n", string(uid), err)
				return
			}
			this.conns[string(uid)] = conn

			// handler read
			go func() {
				for {
					buf := make([]byte, 1024)
					if _, err := conn.Read(buf); err != nil {
						// 出错时断开
						fmt.Printf("连接错误(uid=%s 详细信息: %v)\n", string(uid), err)
						conn.Close()
						return
					}
					rch <- buf
				}
			}()
		}()
	}
}

/*
启动TCP服务器
*/
func ServeTCP(port int, routers map[string]func([]byte) []byte) {
	server := &tcpServer{"0.0.0.0", port, make(map[string]*net.TCPConn), routers}
	server.start()
}
