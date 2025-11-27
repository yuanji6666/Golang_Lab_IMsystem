package main

import (
	"fmt"
	"net"
	"flag"
	"io"
	"os"
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
		flag : 999,
		
	}

	conn, err := net.Dial("tcp",fmt.Sprintf("%s:%d", serverIp,serverPort))
	if err != nil {
		fmt.Println("net.Dial error :", err)
		return nil
	}
	client.conn=conn
	return client
}

func (client *Client) DealResponse(){
	io.Copy(os.Stdout, client.conn)
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
func (client *Client) UpdateName() bool {
	fmt.Println(">>>请输入用户名")
	fmt.Scanln(&client.Name)

	sendMsg := "rename|" +client.Name+"\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn.Write err: ", err)
		return false
	}
	return true
}
func (client *Client) Run(){
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
			client.UpdateName()
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
	go client.DealResponse()
	fmt.Println("---client connect succecd---")

	//start client bussiness
	client.Run()
} 