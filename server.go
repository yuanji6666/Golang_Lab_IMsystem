package main

import (
	"net"
	"fmt"
)
type Server struct {
	Ip string
	Port int
}
//interface : create a server
func NewServer(ip string, port int) *Server {
	server := &Server {
		Ip: ip,
		Port: port,
	}
	return server
}

func (this *Server) Handler(conn net.Conn) {
	// connection bussiness now...
	fmt.Println("connect success")
}

//interface : start a server
func (this *Server) Start() {
	
	//socket listening
	listener,err :=net.Listen("tcp",fmt.Sprintf("%s:%d",this.Ip,this.Port))
	if err!=nil{
		fmt.Println("net.Listen err:",err)
		return
	}
	//close listen socket
	defer listener.Close()

	for{
		//accept
		conn,err := listener.Accept()
		if err!=nil {
			fmt.Println("listener accept err :", err)
			continue;
		}

		go this.Handler(conn)
	}
	 
}

