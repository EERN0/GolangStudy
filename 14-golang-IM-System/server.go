package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
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

	user := NewUser(conn, serv)

	// 用户上线
	user.Online()

	// 监听用户是否活跃的channel
	isAlive := make(chan bool)

	// 接收客户端(用户)发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 { // 读取到的字符个数为0
				user.Offline() // 用户下线
				return
			}

			if err != nil && err != io.EOF {
				fmt.Println("Conn Read err: ", err)
				return
			}

			// 提取用户的消息（去除'\n'）
			msg := string(buf[:n-1])

			// 发给用户的消息处理模块来处理消息，用户的业务(和服务端没关系)
			user.DoMessage(msg)

			// 有用户消息传到服务器，表明该用户是活跃的
			isAlive <- true
		}
	}()

	// 超时强踢功能，添加定时器功能
	// select{} 让当前handler阻塞
	for {
		select {
		case <- isAlive:	// isAlive管道中有值，表明当前用户是活跃的，用来重置定时器
			// isAlive管道中没有值，select阻塞，20s后，下面的管道就会有值，触发下面的case
		case <-time.After(time.Second * 20):	// time.After()是一个channel, 20s之后就有数据了
			// 20s之后，isAlive还没有数据，表明已经超时, 将当前user强制关闭
			user.SendMsg("你被踢了")

			// 销毁用户资源
			close(user.C)

			// 关闭连接
			conn.Close()
			
			// 退出当前go程
			return		// 或者用 runtime.Goexit()
		}
	}
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
