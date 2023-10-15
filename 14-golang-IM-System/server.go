package main

import (
	"fmt"
	"net"
)

type Server struct {
	Ip   string
	Port int
}

// 创建一个server接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:   ip,
		Port: port,
	}
	return server
}

func (this *Server) Handler(conn net.Conn) {
	// ...当前连接的业务
	fmt.Println("链接建立成功")
}

// 启动服务器接口
func (this *Server) Start() { // Server类的成员方法
	// #1 socket listen
	// net.Listen()的操作相当于socket()-bind()-listen()的融合。listener是监听fd
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port)) // address格式：127.0.0.1:8888
	if err != nil {
		fmt.Println("net.Listen error: ", err)
		return
	}
	// close listen socket
	defer listener.Close()

	for {
		// #2 accept
		// conn是通信fd，Accept()是阻塞的
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept error: ", err)
			continue
		}

		// #3 do handler
		go this.Handler(conn)	// 用一个协程来处理，主go程继续执行Accept()
	}
}
