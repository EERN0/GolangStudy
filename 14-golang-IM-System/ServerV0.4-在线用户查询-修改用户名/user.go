package main

import (
	"net"
	"strings"
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

// 给当前User对应的客户端发送消息
func (user *User) SendMsg(msg string) {
	user.conn.Write([]byte(msg))
}

// 用户处理消息的业务
func (user *User) DoMessage(msg string) {
	if msg == "online" { // 查询当前在线用户有哪些
		user.server.mapLock.Lock()
		for _, u := range user.server.OnlineMap {
			onlineMsg := "[" + u.Addr + "]" + u.Name + ":" + "在线...\n"
			user.SendMsg(onlineMsg) // 谁查的发给谁
		}
		user.server.mapLock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename|" { // 修改用户名，命令格式：rename|xxx
		newName := strings.Split(msg, "|")[1]

		// 判断name是否存在
		_, ok := user.server.OnlineMap[newName]
		if ok { // 查询成功
			user.SendMsg("当前用户名已被使用\n")
		} else {
			user.server.mapLock.Lock()
			delete(user.server.OnlineMap, user.Name)
			user.server.OnlineMap[newName] = user
			user.server.mapLock.Unlock()

			user.Name = newName
			user.SendMsg("修改用户名成功:" + user.Name + "\n")
		}
	} else {
		user.server.BroadCast(user, msg)
	}
}
