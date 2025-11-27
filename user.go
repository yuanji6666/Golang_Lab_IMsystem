package main
import (
	"net"
	"strings"
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
//send message to specific user
func (this *User) SendMsg (msg string){
	this.conn.Write([]byte(msg))
}
//do message
func (this *User)DoMessage(msg string){
	if msg=="who"{
		this.server.mapLock.Lock()
		for _,user :=range this.server.OnlineMap {
			onlineMsg := "[" +user.Addr + "]" +user.Name+ ":" + "online..."+"\n"
			this.SendMsg(onlineMsg)
		}
		this.server.mapLock.Unlock()

	}else if len(msg)>7&&msg[:7]=="rename|"{
		newName := msg[7:]
		// if the newName exsited
		_,ok := this.server.OnlineMap[newName]
		if ok {
			this.SendMsg("user name already exists")
		}else {
			this.server.mapLock.Lock()
			delete(this.server.OnlineMap,this.Name)
			this.server.OnlineMap[newName]=this
			this.server.mapLock.Unlock()

			this.Name=newName
			this.SendMsg("success : name changed"+this.Name+ "\n")

		}
	}else if len(msg)>=4&&msg[:3]=="to|"{
		remoteName:= strings.Split(msg,"|")[1] 
		if remoteName ==""{
			this.SendMsg("message format incorrect!")
			return
		}

		remoteUser,ok := this.server.OnlineMap[remoteName]
		if !ok {
			this.SendMsg("user not exist ,please check")
			return 
		}
		content := strings.Split(msg,"|")[2]
		if content ==""{
			this.SendMsg("message is empty,please type in characters")
			return
		}
		remoteUser.SendMsg(this.Name+" send message to you : "+ "["+content+"]")
	}else{
		
		this.server.BroadCast(this,msg)
	}
}

//method : listening to channel and send to the client
func (this * User) ListenMessage(){
	for {
		msg:= <-this.C
		this.conn.Write([]byte(msg+"\n"))
	}
}

