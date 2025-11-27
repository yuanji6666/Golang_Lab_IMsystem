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
	flag int
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

func (client *Client) menu() bool {
	var flag int

	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.重命名")
	fmt.Println("0.结束")

	fmt.Scanln(&flag)

	if flag>= 0 &&flag <= 3{
		client.flag=flag
		return true
	}else{
		fmt.Println("请选择合适范围内的数字")
		return false
	}
}

func (client *Client) Run(){
	client.flag=1
	for client.flag != 0{
		for client.menu() != true {
		}

		switch client.flag{
		case 1:
			fmt.Println("公聊模式启动...")
			break
		case 2:
			fmt.Println("私聊模式启动...")
			break
		case 3:
			fmt.Println("更新用户名选择...")
			break
		}

	}
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
	fmt.Println("---client connect succecd---")

	//start client bussiness
	client.Run()
} 