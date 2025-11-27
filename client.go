package main

import (
	"fmt"
	"net"
	"flag"
)

type Client struct{
	ServerIp string
	ServerPort int
	Name string
	conn net.Conn
}

func NewClient(serverIp string,serverPort int) *Client {
	client := &Client{
		ServerIp : serverIp,
		ServerPort : serverPort,
	}

	conn, err := net.Dial("tcp",fmt.Sprintf("%s:%d", serverIp,serverPort))
	if err != nil {
		fmt.Println("net.Dial error :", err)
		return nil
	}
	client.conn=conn
	return client
}

var serverIp string
var serverPort int
func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置默认服务器ip地址（默认127.0.0.1）")
	flag.IntVar(&serverPort, "port", 8888, "设置服务器端口（默认8888）")
}
func main(){
	
	flag.Parse()

	client := NewClient(serverIp,serverPort)
	if client ==nil {
		fmt.Println("---client connect failed---")
		return	
	}
	fmt.Println("---client connect succecd")

	//start client bussiness
	select{}
} 