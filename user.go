package main

import (
	"net"
	
)

type User struct {
	Name string
	Addr string
	C chan string
	conn net.Conn
}
//interface:create a user
func NewUser(conn net.Conn) *User {
	var userAddr string=conn.RemoteAddr().String()
	user :=&User{
		Name : userAddr,
		Addr : userAddr,
		C : make(chan string),
		conn: conn,
	}
	//start a goroutine which is listening to channel
	go user.ListenMessage()
	return user
}
//method : listening to channel and send to the client
func (this * User) ListenMessage(){
	for {
		msg:= <-this.C
		this.conn.Write([]byte(msg+"\n"))
	}
}

