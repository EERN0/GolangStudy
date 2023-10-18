package main

import (
	"fmt"
	"io"
	"net"
	"sync"
)

/*
 * 服务器类
 */
type Server struct {
	Ip   string
	Port int

	// 在线用户的列表
	OnlineMap map[string]*User
	mapLock   sync.RWMutex

	// 消息广播的channel
	Message chan string
}

// 创建一个server接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

// 广播消息的方法
func (serv *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	serv.Message <- sendMsg
}

// 监听Message广播消息channel的goroutine，一旦Message有消息就发送给全部的在线User
func (serv *Server) ListenMessage() {
	for {
		msg := <-serv.Message

		// 将消息发送给全部的在线User
		serv.mapLock.Lock()
		for _, cli := range serv.OnlineMap {
			cli.C <- msg
		}
		serv.mapLock.Unlock()
	}
}

func (serv *Server) Handler(conn net.Conn) {
	// ...当前连接的业务
	// fmt.Println("链接建立成功")

	user := NewUser(conn)

	// 用户上线，将用户加入到onlineMap中
	serv.mapLock.Lock()
	serv.OnlineMap[user.Name] = user // 多用户操作同一个map，互斥访问
	serv.mapLock.Unlock()

	// 广播当前用户上线消息
	serv.BroadCast(user, "已上线")

	// 接收客户端发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 { // 读取到的字符个数为0
				serv.BroadCast(user, "用户下线")
				return
			}

			if err != nil && err != io.EOF {
				fmt.Println("Conn Read err: ", err)
				return
			}

			// 提取用户的消息（去除'\n'）
			msg := string(buf[:n-1])

			// 将得到的消息进行广播
			serv.BroadCast(user, msg)
		}
	}()

	// 当前handler阻塞
	select {}
}

// 启动服务器接口
func (serv *Server) Start() { // Server类的成员方法
	// #1 socket listen
	// net.Listen()的操作相当于socket()-bind()-listen()的融合。listener是监听fd
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", serv.Ip, serv.Port)) // address格式：127.0.0.1:8888
	if err != nil {
		fmt.Println("net.Listen error: ", err)
		return
	}
	// close listen socket
	defer listener.Close()

	// 启动监听Message的goroutine
	go serv.ListenMessage()

	for {
		// #2 accept
		// conn是通信fd，Accept()是阻塞的
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept error: ", err)
			continue
		}

		// #3 do handler
		go serv.Handler(conn) // 用一个协程goroutine来处理，主go程继续执行Accept()
	}
}
