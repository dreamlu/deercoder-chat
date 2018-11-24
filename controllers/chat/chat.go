package chat

import (
	"deercder-chat/models/chat"
	"github.com/Dreamlu/deercoder-gin/util/lib"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if !strings.Contains(r.URL.Path, "/chat") {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "static/html/chat/index.html")
}

//聊天
func Chat(u *gin.Context) {

	serveHome(u.Writer, u.Request)
}

//聊天ws
func ChatWS(u *gin.Context) {

	go handleMessages()
	WsHander(u.Writer, u.Request)
}

//创建群聊
func DistributeGroup(u *gin.Context) {
	uids := u.PostForm("uids")
	gid, _ := chat.DistributeGroup(uids)
	if gid == "" {
		u.JSON(http.StatusOK, lib.GetMapData(lib.CodeChat, "群聊创建失败"))
		return
	}
	u.JSON(http.StatusOK, map[string]interface{}{"status": lib.CodeSuccess, "msg": "请求成功", "group_id": gid})
}

//拉取群聊所有消息
func GetAllGroupMsg(u *gin.Context) {
	group_id, _ := strconv.ParseInt(u.Query("group_id"), 10, 64)

	msg, err := chat.GetAllGroupMsg(group_id)
	if err != nil {
		u.JSON(http.StatusOK, lib.GetMapData(lib.CodeError, err.Error()))
		return
	}
	var getinfo lib.GetInfoN
	getinfo.Status = lib.CodeSuccess
	getinfo.Msg = lib.MsgSuccess
	getinfo.Data = msg
	u.JSON(http.StatusOK, getinfo)
}

//拉取离线信息
func GetGroupLastMsg(u *gin.Context) {
	group_id, _ := strconv.ParseInt(u.Query("group_id"), 10, 64)
	uid, _ := strconv.ParseInt(u.Query("uid"), 10, 64)

	msg, err := chat.GetGroupLastMsg(group_id, uid)
	if err != nil {
		u.JSON(http.StatusOK, lib.GetMapData(lib.CodeError, err.Error()))
		return
	}
	var getinfo lib.GetInfoN
	getinfo.Status = lib.CodeSuccess
	getinfo.Msg = lib.MsgSuccess
	getinfo.Data = msg
	u.JSON(http.StatusOK, getinfo)
}

// 已读离线信息
func ReadGroupLastMsg(u *gin.Context) {
	group_id, _ := strconv.ParseInt(u.PostForm("group_id"), 10, 64)
	uid, _ := strconv.ParseInt(u.PostForm("uid"), 10, 64)

	ss := chat.ReadGroupLastMsg(group_id, uid)
	u.JSON(http.StatusOK, ss)
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