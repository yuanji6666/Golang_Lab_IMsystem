package main

import (
	"net"
	
)

type User struct {
	Name string
	Addr string
	C chan string
	conn net.Conn
	server *Server
}
//interface:create a user
func NewUser(conn net.Conn,server *Server) *User {
	var userAddr string=conn.RemoteAddr().String()
	user :=&User{
		Name : userAddr,
		Addr : userAddr,
		C : make(chan string),
		conn: conn,
		server :server,
	}
	//start a goroutine which is listening to channel
	go user.ListenMessage()
	return user
}
//online
func (this *User)Online(){
	//add user to onlinemap
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()
	//broadcast new user
	this.server.BroadCast(this,"online...")

}
//offline
func (this *User)Offline(){
	//delete user to onlinemap
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap,this.Name)
	this.server.mapLock.Unlock()
	//broadcast offline user
	this.server.BroadCast(this,"offline...")
	
}
//do message
func (this *User)DoMessage(msg string){
	this.server.BroadCast(this,msg)
}

//method : listening to channel and send to the client
func (this * User) ListenMessage(){
	for {
		msg:= <-this.C
		this.conn.Write([]byte(msg+"\n"))
	}
}

