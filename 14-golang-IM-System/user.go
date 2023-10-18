package main

import (
	"net"
)

/*
 * 用户类
 */
type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn

	server *Server
}

// 创建一个用户API
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: conn,

		server: server,
	}

	// 启动监听当前user channel消息的gorutine
	go user.ListenMessage()

	return user
}

// User类的成员方法
// 监听当前User channel的方法，一旦有消息，就直接发送给对端客户端
func (u *User) ListenMessage() {
	for {
		msg := <-u.C

		u.conn.Write([]byte(msg + "\n"))
	}
}

// 用户的上线业务
func (u *User) Online() {
	// 用户上线，将用户加入到onlineMap中
	u.server.mapLock.Lock()
	u.server.OnlineMap[u.Name] = u
	u.server.OnlineMap[u.Name] = u // 多用户操作同一个map，互斥访问
	u.server.mapLock.Unlock()

	// 广播当前用户上线消息
	u.server.BroadCast(u, "已上线")
}

// 用户的下线业务
func (u *User) Offline() {
	// 用户下线，将用户从onlineMap中删除
	u.server.mapLock.Lock()
	delete(u.server.OnlineMap, u.Name)
	u.server.mapLock.Unlock()

	// 广播当前用户下线消息
	u.server.BroadCast(u, "下线")
}

// 用户处理消息的业务
func (user *User) DoMessage(msg string) {
	user.server.BroadCast(user, msg)
}
