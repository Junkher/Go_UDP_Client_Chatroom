package main

import (
	"context"
	"fmt"
	"net"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//client struct
type client struct {
	ctx         context.Context
	conn        *net.UDPConn
	alive       bool //用来判断用户退出
	inRoom      bool 
	// inroomName  string
	sendMsgs    chan string
	receiveMsgs chan string
	ok          chan bool
	roomList    chan string
	userList    chan string
	// receivedRoomMsgs chan string
	// receivedLobbyMsgs chan string

}

//错误检查函数
func checkError(err error) {
		if err != nil {
			fmt.Println("Error:", err)
		}
}


// 初始化管道,疯狂使用chan :)
func NewClient() *client {
	return &client{
		sendMsgs: make(chan string),
		receiveMsgs: make(chan string),
		ok: make(chan bool),
		roomList: make(chan string),
		userList: make(chan string),
		// receivedRoomMsgs: make(chan string),
		// receivedLobbyMsgs: make(chan string)
	}
}

//应用启动时会调用startup
func (c *client) startup(ctx context.Context) {
	// Perform your setup here
	//初始化
	c.ctx = ctx
	c.alive = false
	c.inRoom = false
}

//应用终止时会调用shutdown
func (c *client) shutdown(ctx context.Context) {
	// Perform your teardown here
	// 退出房间并断开
	if c.inRoom {
		c.conn.Write([]byte("/quit"))
	}
	c.conn.Write([]byte("/disconnect"))
}

func (c *client) Connect(serviceAddr string) {

	if c.alive {
		c.printLobbyMsg("已连接到服务器")
		c.GetRoomList()
		return
	}
	rAddr, err := net.ResolveUDPAddr("udp4", serviceAddr)
	checkError(err)

	fmt.Println(rAddr)
	//Net.UDPdial指定了远程地址，是'connected'类型的连接，因此只需conn.Read
	c.conn, err = net.DialUDP("udp", nil, rAddr) 
  checkError(err)
	fmt.Println("udpconn establish:", serviceAddr)
	defer c.conn.Close()
	c.alive = true
	runtime.EventsEmit(c.ctx, "connect", true)
	go c.GetRoomList()
	go c.freshRoomList()
	go c.handleMsg()
	go c.receive()
  c.send()
	// c.input()
} 
 
func (c *client)  Disconnect() {
	c.sendMsgs <- "/disconnect"
	c.alive = false
	runtime.EventsEmit(c.ctx, "connect", false)
	c.printLobbyMsg( "已离开服务器")
}

func (c *client) GetRoomList() {
	c.sendMsgs <- "/rooms"
	c.printLobbyMsg( "欢迎你，加入房间开始聊天吧 👇")

}

func (c *client) freshRoomList() {
	for  c.alive {
		fmt.Println("等待roomList")
    roomList := <- c.roomList
		fmt.Println("读出roomList")
		runtime.EventsEmit(c.ctx, "roomList", roomList)
		runtime.EventsEmit(c.ctx, "connect", true)
	}
}

func (c *client) EnterRoom(roomName string) {
	c.inRoom = true
	runtime.EventsEmit(c.ctx, "inroom", true)
	runtime.EventsEmit(c.ctx, "roomName", roomName)
	c.sendMsgs <- "/join " + roomName
  // go c.getUserList()
	go c.freshUserList()
}

func (c *client) QuitRoom() {
	c.inRoom = false
	runtime.EventsEmit(c.ctx, "inroom", false)
	c.sendMsgs <- "/quit"
	// c.GetRoomList()
	// runtime.EventsEmit(c.ctx, "connect", c.alive)
}

func (c *client) CreateRoom(roomName string) bool {
	c.sendMsgs <- "/create " + roomName
	ok := <- c.ok
	if ok {
		c.EnterRoom(roomName)
		return true
	} else {
		// c.printLobbyMsg("已存在")
		return false
	}
}

func (c *client) Nick(nickname string) {
	c.sendMsgs <- "/nick " + nickname
}

func (c *client) Broadcast(msg string) {
	c.sendMsgs <- "/msg " + msg
}

func (c *client) getUserList() {
	c.sendMsgs <- "/users"
}

func (c *client) freshUserList() {
	for  c.alive && c.inRoom {
		fmt.Println("等待userList")
    userList := <- c.userList
		fmt.Println("读出userList")
		runtime.EventsEmit(c.ctx, "userList", userList)
		// runtime.EventsEmit(c.ctx, "connect", true)
	}
}


func (c *client) receive() {
	 buf := make([]byte, 1024)
	 for c.alive{
			n, err := c.conn.Read(buf)
			checkError(err)
			fmt.Println("接收信息：", string(buf[0:n]))
			checkError(err)
			c.receiveMsgs <- string(buf[0:n])
			fmt.Println("写进receiveMsgs")
	 }
}


func (c *client) handleMsg() {
	  for c.alive {
			msg := <- c.receiveMsgs
			fmt.Println("读出receiveMsgs")
			// fmt.Println(msg)
			if c.inRoom && msg != "Fresh"  {
				 if strings.HasPrefix(msg, "Users: ") {
				  	msg = strings.TrimPrefix(msg, "Users: ")
						fmt.Println("Users: ", msg)
						fmt.Println("写不进userList")
						c.userList <- msg
						fmt.Println("写进userList")
						continue
				 } 
				 if msg == "FreshUserList" {
					  c.getUserList()
						continue
				 }
				 fmt.Println("Room: ", msg)
				 runtime.EventsEmit(c.ctx, "roomMsg", msg)
			} else {

				if strings.HasPrefix(msg, "Lobby: ") {
					// rooms := strings.Split(msg, ",")
					msg = strings.TrimPrefix(msg, "Lobby: ")
					fmt.Println("Lobby: ", msg)
					fmt.Println("写不进roomList")
					c.roomList <- msg
					fmt.Println("写进roomList")
					continue
				} 
	
				if strings.HasPrefix(msg, "Create: ") {
					msg = strings.TrimPrefix(msg, "Create: ")
					s := strings.Split(msg, " ")
					if s[1] == "ok" {
						// c.EnterRoom(s[0])
						c.ok <- true
					} else {
						// c.printLobbyMsg("该房间名已存在")
						c.ok <- false
					}
				}

				if msg == "Fresh"{
					c.GetRoomList()
					continue
				}
			}
		}
}

func (c *client) printLobbyMsg (msg string) {
	 runtime.EventsEmit(c.ctx, "lobbyMsg", msg)
}


func  (c *client) send() {
	for c.alive {
		msg := <- c.sendMsgs
		fmt.Println(msg)
		_, err := c.conn.Write([]byte(msg))
		fmt.Println("发送了信息: ", msg)
		checkError(err)
	}
}


