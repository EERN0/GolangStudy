package main

func main(){
	server := NewServer("127.0.0.1",8888)	// 服务端的ip和端口
	server.Start()	// 启动服务器
}