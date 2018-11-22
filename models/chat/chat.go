package chat

import (
	"errors"
	"fmt"
	"github.com/Dreamlu/deercoder-gin"
	"github.com/Dreamlu/deercoder-gin/util/lib"
	"strconv"
	"strings"
	"time"
)

// Define our message object,teacher message model
type Message struct {
	ID          int64         `json:"id"`           //群组消息id
	GroupId     string        `json:"group_id"`     //组id
	FromUid     int64         `json:"from_uid"`     //来自用户id
	Headimg     string        `json:"headimg"`      //头像
	Username    string        `json:"username"`     //用户名
	Content     string        `json:"content"`      //消息内容
	Flag        int64         `json:"flag"`         //0老师,1学生
	ContentType string        `json:"content_type"` //前台用
	CreateTime  deercoder.JsonTime `json:"create_time"`  //创建时间
	//SendFrom	string `json:"send_from"`
}

/*群聊发送模型*/
type GroupMsg struct {
	ID         int64         `json:"id"`
	GroupID    string        `json:"group_id"`    //群聊id
	Content    int64         `json:"content"`     //消息内容
	FromUid    int64         `json:"from_uid"`    //由谁发送
	Flag       int64         `json:"flag"`        //0老师,1学生
	CreateTime deercoder.JsonTime `json:"create_time"` //创建时间
}

///*聊天内容*///暂时废弃
//type Msg struct {
//	ID         int64         `json:"id"`
//	Content    string        `json:"content"`
//	CreateTime util.JsonTime `json:"create_time"`
//}

/*聊天配置*/
type GroupLastMsg struct {
	ID             int64  `json:"id"`
	GroupID        string `json:"group_id"` //群聊id
	Uid            int64  `json:"uid"`
	Flag           int64  `json:"flag"` //0老师,1学生
	LastGroupMsgId int64  `json:"last_group_msg_id"`
}

/*群组id极其成员id*/
type GroupUsers struct {
	ID      int64  `json:"id"`
	GroupId string `json:"group_id"`
	Uid     int64  `json:"uid"`
	Flag    int64  `json:"flag"`
}

////群组,删除
//func DeleteGroup(){
//
//}

//建立群组,未来扩展
//返回群组id
func DistributeGroup(studentids, teacherids string) (groupId int64, err error) {

	if studentids == "" && teacherids == "" {
		return 0, nil
	}

	sids := strings.Split(studentids, ",")
	//唯一群号
	//gidstr := time.Now().Format("20060102150405")
	//gidstr = string([]byte(gidstr)[2:])
	groupId = time.Now().UnixNano() //纳秒 //strconv.ParseInt(gidstr, 10, 64)
	gidstr := strconv.FormatInt(groupId, 10)
	sql := "insert `group_users`(group_id,uid,flag) value"
	for _, v := range sids { //学生
		if v == "" {
			continue
		}
		sql += "(" + gidstr + "," + v + ",1),"
	}

	tids := strings.Split(teacherids, ",")
	for _, v := range tids { //老师
		if v == "" {
			continue
		}
		sql += "(" + gidstr + "," + v + ",0),"
	}

	sql = string([]byte(sql)[:len(sql)-1])

	dba := deercoder.DB.Exec(sql)
	num := dba.RowsAffected

	if num == 0 {
		return 0, dba.Error
	}

	return groupId, nil
}

//群聊消息,创建
func CreateGroupMsg(id int64, group_id string, from_uid, flag int64, content, content_type string) (err error) {

	//需要id,用来每次聊天生成的id作为聊天记录id,以便群离线消息记录该id
	sql := "insert `group_msg`(id, group_id,content,from_uid, flag, content_type) value(?,?,?,?,?,?)"
	dba := deercoder.DB.Exec(sql, id, group_id, content, from_uid, flag, content_type)

	if dba.Error != nil {
		return dba.Error
	}
	return nil
}

//群离线消息记录
//记录用户离线时,最后显示的消息id
func CreateGroupLastMsg(group_id string, uid, flag, last_group_msg_id int64) (err error) {
	if group_id == "" {
		return errors.New("用户gid不存在")
	}
	sql := "insert `group_last_msg`(group_id, uid, flag, last_group_msg_id) value(?,?,?,?)"
	dba := deercoder.DB.Exec(sql, group_id, uid, flag, last_group_msg_id)

	if dba.Error != nil {
		return dba.Error
	}
	return nil
}

//拉取用户离线消息
func GetAllGroupMsg(group_id int64) ([]Message, error) {

	//拉取该群聊的所有消息
	sql := `select *
	from group_msg
	where group_id=?`

	var msg []Message

	deercoder.DB.Raw(sql, group_id).Scan(&msg)

	if len(msg) == 0 {

		return msg, errors.New("暂无离线消息")
	}
	for k, v := range msg { //查询对应的头像,用户名等信息
		sql := "select username,headimg from `teacher` where id = ?"
		switch v.Flag {
		case 0: //老师
		case 1: //学生
			sql = strings.Replace(sql, "teacher", "student", -1)
		}
		deercoder.DB.Raw(sql, v.FromUid).Scan(&msg[k])
	}

	return msg, nil
}

//拉取用户离线消息
func GetGroupLastMsg(group_id, uid, flag int64) ([]Message, error) {

	//1.找出群聊group_id中对应的最小的未读记录id
	var value deercoder.Value
	sql2 := "select min(last_group_msg_id) as value from group_last_msg where is_read=0 and group_id=? and uid=? and flag=?"
	deercoder.DB.Raw(sql2, group_id, uid, flag).Scan(&value)

	if value.Value == "" {
		return nil, errors.New("暂无离线消息")
	}
	//2.拉取离线后的该群聊的所有消息
	sql := `select *
	from group_msg
	where group_id=? and id >= ?`

	var msg []Message

	deercoder.DB.Raw(sql, group_id, value.Value).Scan(&msg)

	if len(msg) == 0 {

		return msg, errors.New("暂无离线消息")
	}
	for k, v := range msg { //查询对应的头像,用户名等信息
		sql := "select username,headimg from `teacher` where id = ?"
		switch v.Flag {
		case 0: //老师
		case 1: //学生
			sql = strings.Replace(sql, "teacher", "student", -1)
		}
		deercoder.DB.Raw(sql, v.FromUid).Scan(&msg[k])
	}

	return msg, nil
}

//已读消息
func ReadGroupLastMsg(group_id, uid, flag int64) interface{} {

	var info interface{}
	sql2 := "update `group_last_msg` set is_read=1 where is_read=0 and group_id=? and uid=? and flag=?"
	dba := deercoder.DB.Exec(sql2, group_id, uid, flag)
	num := dba.RowsAffected
	if dba.Error != nil {
		info = lib.GetSqlError(dba.Error.Error())
	} else if num == 0 && dba.Error == nil {
		info = lib.MapExistOrNo
	} else {
		info = lib.MapUpdate
	}
	return info
}

// 群发消息,对方默认未读
// flag 0老师, 1学生
// group_ids,逗号分割,群聊id
// send_uids老师或学生id,和群聊id一一对应
func MassMessage(group_ids, send_uids, from_uid, content, flag string) interface{} {

	if group_ids == "" {
		return lib.GetMapData(lib.CodeChat, "group_ids不能为空")
	}
	sql := "insert `group_msg`(id, group_id, content, from_uid, flag) value"
	sql2 := "insert `group_last_msg`(group_id, uid, flag, last_group_msg_id) value"
	id := time.Now().UnixNano()
	last_msg_id := id
	gids := strings.Split(group_ids, ",")
	uids := strings.Split(send_uids, ",")
	for k, v := range gids {
		sql += fmt.Sprintf("(%d,'%s','%s','%s',%d),", id, v, content, from_uid, 0) //这里肯定是老师群发,flag直接为０
		sql2 += fmt.Sprintf("('%s','%s','%s',%d),", v, uids[k], flag, last_msg_id)
		id = time.Now().UnixNano()
	}

	sql = string([]byte(sql)[:len(sql)-1])    //去,
	sql2 = string([]byte(sql2)[:len(sql2)-1]) //去,
	deercoder.DB.Exec(sql)
	dba := deercoder.DB.Exec(sql2) //创建存储群聊消息

	if dba.Error != nil {
		return lib.GetSqlError(dba.Error.Error())
	}

	return lib.MapCreate
}

//查找群聊中所有用户
func GetChatUsers(group_id string) []GroupUsers {

	var gusers []GroupUsers
	deercoder.DB.Raw("select id,group_id,uid,flag from `group_users` where group_id=?", group_id).Scan(&gusers)
	return gusers
}