package chat

import (
	"context"
	"deercoder-chat/api-gateway/conf"
	"deercoder-chat/chat-srv/models/chat"
	"deercoder-chat/chat-srv/proto"
	"github.com/dreamlu/deercoder-gin/util/lib"
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
	cli = proto.NewStreamerService("deercoder-chat.chat", rpcClient)
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
	chatClient proto.ChatService
)

func init() {
	chatClient = proto.NewChatService(conf.ChatSrv, client.DefaultClient)
}

// 创建群聊/好友
func DistributeGroup(u *gin.Context) {

	uids := u.PostForm("uids")

	// rpc service
	res, err := chatClient.DistributeGroup(context.TODO(), &proto.UidS{
		Uids: uids,
	})

	if err != nil {
		u.JSON(http.StatusInternalServerError, lib.GetMapDataError(err.Error()))
		return
	}

	u.JSON(http.StatusOK, lib.GetMapDataSuccess(map[string]interface{}{"group_id": res.Message.GroupId}))
}

// 拉取群聊所有消息
func GetAllGroupMsg(u *gin.Context) {
	group_id := u.Query("group_id")

	// rpc service
	res, err := chatClient.GetAllGroupMsg(context.TODO(), &proto.Request{
		Message: &proto.Message{
			GroupId: group_id,
		},
	})

	if err != nil {
		u.JSON(http.StatusInternalServerError, lib.GetMapDataError(err.Error()))
		return
	}

	u.JSON(http.StatusOK, lib.GetMapDataSuccess(res.Message))
}

// 拉取离线信息
func GetGroupLastMsg(u *gin.Context) {
	group_id := u.Query("group_id")
	uid, _ := strconv.ParseInt(u.Query("uid"), 10, 64)

	// rpc service
	res, err := chatClient.GetGroupLastMsg(context.TODO(), &proto.Request{
		Message: &proto.Message{
			GroupId: group_id,
			FromUid: uid,
		},
	})

	if err != nil {
		u.JSON(http.StatusInternalServerError, lib.GetMapDataError(err.Error()))
		return
	}

	u.JSON(http.StatusOK, lib.GetMapDataSuccess(res))
}

// 已读离线信息
func ReadGroupLastMsg(u *gin.Context) {
	group_id := u.PostForm("group_id")
	uid, _ := strconv.ParseInt(u.PostForm("uid"), 10, 64)

	// rpc service
	res, err := chatClient.ReadGroupLastMsg(context.TODO(), &proto.Request{
		Message: &proto.Message{
			GroupId: group_id,
			FromUid: uid,
		},
	})

	if err != nil {
		u.JSON(http.StatusInternalServerError, lib.GetMapDataError(err.Error()))
		return
	}
	u.JSON(http.StatusOK, lib.GetMapDataSuccess(res))
}

// 群发消息
func MassMessage(u *gin.Context) {

	group_ids := u.PostForm("group_ids")
	send_uids := u.PostForm("send_uids")
	from_uid := u.PostForm("from_uid")
	content := u.PostForm("content")
	ss := chat.MassMessage(group_ids, send_uids, from_uid, content)
	u.JSON(http.StatusOK, ss)
}
