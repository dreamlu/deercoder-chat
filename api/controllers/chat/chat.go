package chat

import (
	"context"
	"deercoder-chat/api/conf"
	"deercoder-chat/chat-srv/models/chat"
	"deercoder-chat/chat-srv/proto"
	"github.com/dreamlu/go-tool/tool/result"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
	"log"
	"net/http"
	"strconv"
	"time"
)

//type Streamer struct{}
//
var (
	cli proto.StreamerService
)

func init() {
	// New RPC client
	rpcClient := client.NewClient(client.RequestTimeout(time.Second * 120))
	cli = proto.NewStreamerService(conf.ChatSrv, rpcClient)
}

//聊天ws
func ChatWS(u *gin.Context) {

	go handleMessages()

	ws, err := upgrader.Upgrade(u.Writer, u.Request, nil)
	if err != nil {
		log.Fatal(err)
	}
	WsHander(cli, ws)
}

// chat service

var (
	ChatClient proto.ChatService
)

func init() {
	ChatClient = proto.NewChatService(conf.ChatSrv, client.DefaultClient)
}

// 创建群聊/好友
func DistributeGroup(u *gin.Context) {

	uids := u.PostForm("uids")

	// rpc service
	res, err := ChatClient.DistributeGroup(context.TODO(), &proto.UidS{
		Uids: uids,
	})

	if err != nil {
		u.JSON(http.StatusOK, result.GetError(err.Error()))
		return
	}

	u.JSON(http.StatusOK, result.GetSuccess(map[string]interface{}{"group_id": res.Message.GroupId}))
}

// 拉取群聊所有消息
func GetAllGroupMsg(u *gin.Context) {
	group_id := u.Query("group_id")

	// rpc service
	res, err := ChatClient.GetAllGroupMsg(context.TODO(), &proto.Request{
		Message: &proto.Message{
			GroupId: group_id,
		},
	})

	if err != nil {
		u.JSON(http.StatusOK, result.GetError(err.Error()))
		return
	}

	u.JSON(http.StatusOK, result.GetSuccess(res.Message))
}

// 拉取离线信息
func GetGroupLastMsg(u *gin.Context) {
	group_id := u.Query("group_id")
	uid, _ := strconv.ParseInt(u.Query("uid"), 10, 64)

	// rpc service
	res, err := ChatClient.GetGroupLastMsg(context.TODO(), &proto.Request{
		Message: &proto.Message{
			GroupId: group_id,
			FromUid: uid,
		},
	})

	if err != nil {
		u.JSON(http.StatusOK, result.GetError(err.Error()))
		return
	}

	u.JSON(http.StatusOK, result.GetSuccess(res))
}

// 已读离线信息
func ReadGroupLastMsg(u *gin.Context) {
	group_id := u.PostForm("group_id")
	uid, _ := strconv.ParseInt(u.PostForm("uid"), 10, 64)

	// rpc service
	res, err := ChatClient.ReadGroupLastMsg(context.TODO(), &proto.Request{
		Message: &proto.Message{
			GroupId: group_id,
			FromUid: uid,
		},
	})

	if err != nil {
		u.JSON(http.StatusOK, result.GetError(err.Error()))
		return
	}
	u.JSON(http.StatusOK, result.GetSuccess(res))
}

// 获取好友列表
func GetUserList(u *gin.Context) {
	uid, _ := strconv.ParseInt(u.Query("uid"), 10, 64)

	// rpc service
	res, err := ChatClient.GetUserList(context.TODO(), &proto.ChatUser{
		Id: uid,
	})

	if err != nil {
		u.JSON(http.StatusOK, result.GetError(err.Error()))
		return
	}
	u.JSON(http.StatusOK, result.GetSuccess(res))
}

// 搜索获取好友列表
func GetUserSearchList(u *gin.Context) {
	uid, _ := strconv.ParseInt(u.Query("uid"), 10, 64)
	name := u.Query("name")

	// rpc service
	res, err := ChatClient.GetUserSearchList(context.TODO(), &proto.ChatUserSearch{
		Id:   uid,
		Name: name,
	})

	if err != nil {
		u.JSON(http.StatusOK, result.GetError(err.Error()))
		return
	}
	u.JSON(http.StatusOK, result.GetSuccess(res))
}

// 获取群聊中用户列表
func GetGroupUser(u *gin.Context) {
	gid := u.PostForm("group_id")

	// rpc service
	res, err := ChatClient.GetGroupUser(context.TODO(), &proto.GroupUser{
		GroupId: gid,
	})

	if err != nil {
		u.JSON(http.StatusOK, result.GetError(err.Error()))
		return
	}
	u.JSON(http.StatusOK, result.GetSuccess(res))
}

//// 创建聊天记录记录
//func CreateGroupMsg(u *gin.Context) {
//	gid := u.PostForm("group_id")
//
//	// rpc service
//	res, err := chatClient.GetGroupUser(context.TODO(), &proto.GroupUser{
//		GroupId: gid,
//	})
//
//	if err != nil {
//		u.JSON(http.StatusOK, result.GetError(err.Error()))
//		return
//	}
//	u.JSON(http.StatusOK, result.GetSuccess(res))
//}

// 群发消息
func MassMessage(u *gin.Context) {

	group_ids := u.PostForm("group_ids")
	send_uids := u.PostForm("send_uids")
	from_uid := u.PostForm("from_uid")
	content := u.PostForm("content")
	ss := chat.MassMessage(group_ids, send_uids, from_uid, content)
	u.JSON(http.StatusOK, ss)
}
