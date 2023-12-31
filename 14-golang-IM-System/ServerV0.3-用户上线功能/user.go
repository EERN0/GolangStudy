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
}

// 创建一个用户API
func NewUser(conn net.Conn) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: conn,
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