package main

import (
	"net"
	"fmt"
	"sync"
)
type Server struct {
	Ip string
	Port int
	//list of online users
	OnlineMap map[string]*User
	mapLock sync.RWMutex

	//channel of message broadcast 
	Message chan string
}
//interface : create a server
func NewServer(ip string, port int) *Server {
	server := &Server {
		Ip: ip,
		Port: port,
		OnlineMap : make(map[string]*User),
		Message : make(chan string),
	}
	return server
}

//listen to broadcast message and send to channel of users

func (this *Server) ListenMessager() {
	for {
		msg:=<-this.Message
		
		//send to all online users
		this.mapLock.Lock()
		for _,cli :=range this.OnlineMap {
			cli.C<- msg
		}
		this.mapLock.Unlock()
	}
}

func (this *Server) BroadCast(user *User,msg string){
	sendMsg:= "["+user.Addr+"]"+user.Name+":"+msg
	this.Message <-sendMsg
}

func (this *Server) Handler(conn net.Conn) {
	// connection bussiness now...
	user := NewUser(conn)
	//add user to onlinemap
	this.mapLock.Lock()
	this.OnlineMap[user.Name] = user
	this.mapLock.Unlock()
	//broadcast new user
	this.BroadCast(user,"online...")
	//
	select{}
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

	go this.ListenMessager()

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

