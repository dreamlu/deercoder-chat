package chat

import (
	"deercoder-chat/chat-srv/proto"
	"errors"
	"fmt"
	"github.com/dreamlu/go-tool"
	"github.com/dreamlu/go-tool/util/lib"
	"github.com/dreamlu/go.uuid"
	"strings"
	"time"
)

// Define our message object,teacher message model
type Message struct {
	UUID        string             `json:"uuid"`         //群组消息id
	GroupId     string             `json:"group_id"`     //组id
	FromUid     int64              `json:"from_uid"`     //来自用户id
	Headimg     string             `json:"headimg"`      //头像
	Name        string             `json:"username"`     //用户名
	Content     string             `json:"content"`      //消息内容
	ContentType string             `json:"content_type"` //消息类型,文字'text',图片'img',2....
	CreateTime  der.JsonTime `json:"create_time"`  //创建时间
}

// 群聊发送模型
type GroupMsg struct {
	ID          int64              `json:"id"`
	GroupID     string             `json:"group_id"`     //群聊id
	Content     int64              `json:"content"`      //消息内容
	FromUid     int64              `json:"from_uid"`     //由谁发送
	ContentType string             `json:"content_type"` //消息类型,文本'text'
	CreateTime  der.JsonTime `json:"create_time"`  //创建时间
}

// 群聊最后记录
type GroupLastMsg struct {
	ID             int64  `json:"id"`
	GroupID        string `json:"group_id"` //群聊id
	Uid            int64  `json:"uid"`
	LastGroupMsgId int64  `json:"last_group_msg_id"`
}

// 群组id极其成员id
type GroupUser struct {
	ID      int64  `json:"id"`
	GroupId string `json:"group_id"`
	Uid     int64  `json:"uid"`
}

////群组,删除
//func DeleteGroup(){
//
//}

// 添加好友
// 建立群组,未来扩展
// 返回群组id
func DistributeGroup(uids string) (groupId string, err error) {

	if uids == "" {
		return "", nil
	}

	// 判断是否已经是好友
	groupUserSql := `select (count(*) > 1) as num 
					from group_users a 
					where uid in (`+uids+`) 
					group by a.group_id
					having num = 1`
	res := der.ValidateSQL(groupUserSql)
	if res == lib.MapValSuccess {
		return "", errors.New("好友已存在")
	}


	userids := strings.Split(uids, ",")
	//唯一群id
	groupId = uuid.NewV1().String()
	sql := "insert `group_users`(group_id, uid, create_time) value"
	for _, v := range userids {
		if v == "" {
			continue
		}
		sql += "('" + groupId + "'," + v + ",'" + time.Now().Format("2006-01-02 15:04:05") + "'),"
	}
	sql = string([]byte(sql)[:len(sql)-1])

	dba := der.DB.Exec(sql)
	num := dba.RowsAffected

	if num == 0 {
		return "", dba.Error
	}

	return groupId, nil
}

// 聊天记录
// 群聊消息,创建
func CreateGroupMsg(msg *proto.Message) (err error) {

	//需要id,用来每次聊天生成的id作为聊天记录id,以便群离线消息记录该id
	sql := "insert `group_msg`(uuid, group_id, content, from_uid, content_type) value(?,?,?,?,?)"
	dba := der.DB.Exec(sql, msg.Uuid, msg.GroupId, msg.Content, msg.FromUid, msg.ContentType)

	if dba.Error != nil {
		return dba.Error
	}
	return nil
}

// 群离线消息记录
// 记录用户离线时,最后显示的消息id
func CreateGroupLastMsg(group_id string, uid int64, last_group_msg_uuid string) (err error) {
	if group_id == "" {
		return errors.New("用户gid不存在")
	}
	sql := "insert `group_last_msg`(group_id, uid, last_group_msg_uuid) value(?,?,?)"
	dba := der.DB.Exec(sql, group_id, uid, last_group_msg_uuid)

	if dba.Error != nil {
		return dba.Error
	}
	return nil
}

// 拉取群聊消息(所有)
// 待优化
func GetAllGroupMsg(group_id string) ([]*proto.Message, error) {

	//拉取该群聊的所有消息
	sql := fmt.Sprintf("select %s from group_msg where group_id=?", der.GetColSql(GroupMsg{}))
	var msg []*proto.Message

	der.DB.Raw(sql, group_id).Scan(&msg)

	if len(msg) == 0 || msg[0].GroupId == "" {
		return msg, errors.New("暂无离线消息")
	}
	sql = "select name, headimg from `user` where id = ?"
	for k, v := range msg { //查询对应的头像,用户名等信息

		der.DB.Raw(sql, v.FromUid).Scan(&msg[k])
	}

	return msg, nil
}

// 拉取用户离线消息
func GetGroupLastMsg(group_id string, uid int64) ([]*proto.Message, error) {

	//1.找出群聊group_id中对应的最小的未读记录id
	var value der.Value
	sql2 := "select min(last_group_msg_uuid) as value from group_last_msg where is_read=0 and group_id=? and uid=?"
	der.DB.Raw(sql2, group_id, uid).Scan(&value)

	if value.Value == "" {
		return nil, errors.New("暂无离线消息")
	}
	//2.拉取离线后的该群聊的所有消息
	sql := fmt.Sprintf("select %s from group_msg where group_id=? and id >= (select id from group_msg where uuid = ?)", der.GetColSql(GroupMsg{}))

	var msg []*proto.Message

	der.DB.Raw(sql, group_id, value.Value).Scan(&msg)

	if len(msg) == 0 {

		return msg, errors.New("暂无离线消息")
	}
	sql = "select name, headimg from `user` where id = ?"
	for k, v := range msg { //查询对应的头像,用户名等信息

		der.DB.Raw(sql, v.FromUid).Scan(&msg[k])
	}

	return msg, nil
}

// 已读消息
func ReadGroupLastMsg(group_id string, uid int64) (lib.MapData, error) {

	var info lib.MapData
	sql2 := "update `group_last_msg` set is_read=1 where is_read=0 and group_id=? and uid=?"
	dba := der.DB.Exec(sql2, group_id, uid)
	num := dba.RowsAffected
	if dba.Error != nil {
		return info, dba.Error
	} else if num == 0 && dba.Error == nil {
		return lib.MapNoResult, errors.New(lib.MsgNoResult)
	} else {
		return lib.MapUpdate, nil
	}
}

// 群发消息,对方默认未读
// flag 0老师, 1学生
// group_ids,逗号分割,群聊id
// send_uids老师或学生id,和群聊id一一对应
func MassMessage(group_ids, send_uids, from_uid, content string) interface{} {

	if group_ids == "" {
		return lib.GetMapData(lib.CodeChat, "group_ids不能为空")
	}
	sql := "insert `group_msg`(uuid, group_id, content, from_uid) value"
	sql2 := "insert `group_last_msg`(group_id, uid, last_group_msg_uuid) value"
	uuidS := uuid.NewV1().String()
	gids := strings.Split(group_ids, ",")
	uids := strings.Split(send_uids, ",")
	for k, v := range gids {
		sql += fmt.Sprintf("(%s,'%s','%s','%s'),", uuidS, v, content, from_uid) //这里肯定是老师群发,flag直接为０
		sql2 += fmt.Sprintf("('%s','%s','%s'),", v, uids[k], uuidS)
		uuidS = uuid.NewV1().String()
	}

	sql = string([]byte(sql)[:len(sql)-1])    //去,
	sql2 = string([]byte(sql2)[:len(sql2)-1]) //去,
	der.DB.Exec(sql)
	dba := der.DB.Exec(sql2) //创建存储群聊消息

	if dba.Error != nil {
		return lib.GetSqlError(dba.Error.Error())
	}

	return lib.MapCreate
}

// 查找用户好友列表
// 排除自己
func GetUserList(uid int64) (users []*proto.ChatUser, err error) {

	sql := `select a.id,name,headimg,introduce,createtime,b.group_id 
			from ` + "`user`" + ` a inner join ` + "`group_users`" + ` b on a.id = b.uid 
			where b.group_id in (select group_id from ` + "`group_users`" + ` where uid = ?) 
			and a.id != ?`
	dba := der.DB.Raw(sql, uid, uid).Scan(&users)
	return users, dba.Error
}

// 搜索用户好友列表中指定用户
// 排除自己
func GetUserSearchList(uid int64, name string) (users []*proto.ChatUser, err error) {

	sql := `select a.id,name,headimg,introduce,createtime,b.group_id 
			from ` + "`user`" + ` a inner join ` + "`group_users`" + ` b on a.id = b.uid 
			where b.group_id in (select group_id from ` + "`group_users`" + ` where uid = ? and name like '%` + name + `%') 
			and a.id != ?`
	dba := der.DB.Raw(sql, uid, uid).Scan(&users)
	return users, dba.Error
}

// 查找群聊中所有用户
// 包含自己
func GetGroupUser(group_id string, gusers []*proto.GroupUser) error {

	dba := der.DB.Raw("select id,group_id,uid from `group_users` where group_id = ?", group_id).Scan(&gusers)
	return dba.Error
}
