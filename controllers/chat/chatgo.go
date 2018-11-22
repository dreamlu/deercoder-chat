package chat

import (
	"deercder-chat/models/chat"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

//客户端
type Client struct {
	GroupID string //标识客户端
	UID     int64	//结合flag，唯一标识用户id
	Flag	int64
	Conn    *websocket.Conn
}

var clients []*Client //客户端队列,指针同步同一个client data
//var clients = make(map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan chat.Message) // broadcast channel

// Configure the upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//开启不同进程代表对应的客户端通信
func WsHander(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()

	// Register our new client
	//clients[ws] = true
	//注册客户端连接
	var ct Client
	ct.Conn = ws
	//放入连接队列
	clients = append(clients, &ct)

	//消息读取,每个客户端数据
	for {
		var msg chat.Message
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		fmt.Println("聊天测试", msg)
		if err != nil {
			log.Printf("error: %v", err)
			//delete(clients, ws) //删除对应连接
			for _, v := range clients { //删除对应连接,emm...暂时先遍历删除～
				//fmt.Println(v)
				if v.Conn == ws {
					////fmt.Println("删除,用户gid:",ct.GroupID)
					//clients = append(clients[:k], clients[k+1:]...)//交给广播处理
					//
					////记录该用户最后读的消息id,用户进程中处理,这里gid已经为0
					//chat.CreateGroupLastMsg(msg.GroupId, msg.FromUid, msg.Flag, msg.ID)
					break
				}
			}
			break
		}
		//fmt.Println(msg)
		msg.ID = time.Now().UnixNano()
		ct.GroupID = msg.GroupId //客户端唯一标识
		ct.UID = msg.FromUid
		ct.Flag = msg.Flag
		//fmt.Println("用户flag",msg.Flag)
		// Send the newly received message to the broadcast channel
		broadcast <- msg

		//send broadcast, then save the message
		chat.CreateGroupMsg(msg.ID, msg.GroupId, msg.FromUid, msg.Flag, msg.Content, msg.ContentType)
	}
}

//消息写入,消息推送(不通进程代表各自客户端的写入进程)
func handleMessages() {
	for {
		msg := <-broadcast //广播
		//获得广播内容,筛选获得group_id,并将内容分发给指定的客户
		for k, client := range clients {
			// send message to every specified client, hehe~
			if client.GroupID != msg.GroupId { // must same group_id
				continue
			}
			// next have same group_id
			err := client.Conn.WriteJSON(msg)
			if err != nil { //当连接挂了
				//fmt.Println("客户:",client,"聊天记录写入失败")
				log.Printf("error: %v", err)
				client.Conn.Close()
				clients = append(clients[:k], clients[k+1:]...)
				////记录该用户最后读的消息id,广播中处理,待优化
				//chat.CreateGroupLastMsg(msg.GroupId,msg.FromUid,msg.Flag,msg.ID)
				continue
			}
		}
		//连接该断的也断了
		//进行用户在线检测
		gusers := chat.GetChatUsers(msg.GroupId)
	into:for _, v2 := range clients {
		if v2.GroupID == msg.GroupId { //在线用户
			for k,v := range gusers {
				if v2.UID == v.Uid && v2.Flag == v.Flag{
					gusers = append(gusers[:k], gusers[k+1:]...) //去除在线用户
					goto into
				}
			}
		}
	}
		//剩下的为群聊离线用户
		//记录离线消息
		for _, v := range gusers{
			chat.CreateGroupLastMsg(msg.GroupId,v.Uid,v.Flag,msg.ID)
		}
	}
}
