// api地址等信息
var api = "localhost:8006";
var myApi = "http://" + api + "/api/v1";
var myWsApi = "ws://" + api + "/api/v1";

// 登陆
function login() {
    const name = $("#name").val();
    const password = $("#password").val();

    $.ajax({
        url: myApi + "/login/login",
        method: "POST",
        data: {
            name: name,
            password: password
        },
        success: function (res) {
            //console.log(res)
            if (res.status === 200) {
                // 储存用户id
                //$.data(myObj, "uid", res.data.id)
                Cookies.set("uid", res.data.id)

                location.href = "index.html"
            } else if (res.status === 211) {
                confirm(res.msg)
            }
        }
    })

}

// 个人信息
function myInfo() {
    // 用户id
    const uid = Cookies.get("uid")

    // 请求数据
    $.ajax({
        url: myApi + "/user/id/" + uid,
        method: "GET",
        data: {
            uid: uid,
        },
        success: function (res) {
            if (res.status === 200) {
                // 加载个人数据
                res.data.headimg = myApi + "/" + res.data.headimg;
                $("#profile").html($("#myInfo").render(res.data))

                // 存储个人数据
                Cookies.set("myInfo", res.data)
            }
        }
    })
}

// 好友列表
function userList() {
    // 用户id
    const uid = Cookies.get("uid")

    // 请求数据
    $.ajax({
        url: myApi + "/chat/getUserList",
        method: "GET",
        data: {
            uid: uid,
        },
        success: function (res) {
            if (res.status === 200) {
                // 修改数据
                // 通过模板引擎渲染数据

                // 我一个后端写起前端,还用起看起来牛逼的用法
                // 我猖狂了,嗯,对,我猖狂了
                res.data.userList.forEach(
                    function (val) {
                        //console.log(val);
                        // 看起来是深拷贝, 借助引用地址的样子
                        val.headimg = myApi + "/" + val.headimg;
                    }
                )
                // 渲染数据
                $("#contact").html($("#userList").render(res.data.userList))

            } else {
                confirm(res.msg)
            }
        }
    })
}

//=========================================================================================

// class wsMessage {
//     constructor(name, headimg, content, group_id, from_uid, content_type){
//         this.name = name;
//         this.headimg = headimg;
//         this.content = content;
//         this.group_id = group_id;
//         this.from_uid = from_uid;
//         this.content_type = content_type;
//     }
//
//     static getMessage(){
//         return {
//
//         }
//     }
// }

// websocket
var ws;

// 聊天内容
var msgData;

// 点击好友列表
// 更新好友信息
// 点击聊天事件
$("#contact").on("click", ".contact", function () {

    const name = $(this).find(".name").text();
    const headimg = $(this).find(".headimg").attr("src");
    // 群聊id
    const groupId = $(this).find(".groupId").text();


    // 好友信息
    // 聊天界面用户头像名称等信息
    const arrayData = {
        'name': name,
        'headimg': headimg,
        'groupId': groupId
    };
    $("#groupUser").html($("#groupUserInfo").render(arrayData));

    // 清除界面数据
    $('.messages ul').empty();

    // 拉取该群聊所有消息
    getAllMsg(groupId);

    // 建立在线websocket即时通讯
    // 开始聊天
    if ("WebSocket" in window) {
        ws = new WebSocket(myWsApi + "/chat/chatWs");

        // 建立ws连接
        ws.onopen = function () {
            // 拉取离线数据

        };

        // 接收数据
        ws.onmessage = function (res) {
            // console.log("[receive data]: " + res.data);
            // 返回的数据json转换
            const data = JSON.parse(res.data);
            // 设置聊天内容样式
            setMsgContent(data);
        };

        // 连接关闭
        ws.onclose = function () {
            // 关闭 websocket
            alert("连接已关闭...");
        };
    } else {
        // 浏览器不支持 WebSocket
        alert("您的浏览器不支持 WebSocket!");
    }

});

// 搜索好友
// 回车搜索
$("#search").on("keydown", function (e) {
    if (e.which === 13) {
        // 获取内容
        const content = $("#search input").val();
        if ($.trim(content) === '') {
            return false;
        }

        // 请求数据
        // 用户id
        const uid = Cookies.get("uid");

        // 请求数据
        $.ajax({
            url: myApi + "/chat/getUserSearchList",
            method: "GET",
            data: {
                uid: uid,
                name: content,
            },
            success: function (res) {
                if (res.status === 200) {
                    // 修改数据
                    // 通过模板引擎渲染数据

                    res.data.userList.forEach(
                        function (val) {
                            val.headimg = myApi + "/" + val.headimg;
                        }
                    );

                    if (res.data.length === 0) {
                        // 清除数据
                        $("#contact").empty();
                    }

                    // 渲染数据
                    $("#contact").html($("#userList").render(res.data.userList))

                } else {
                    confirm(res.msg)
                }
            }
        })
    }
});

//===========================================================

// 拉取该群聊所有聊天记录
function getAllMsg(group_id) {

    $.ajax({
        url: myApi + "/chat/allMsg/",
        method: "GET",
        data: {
            group_id: group_id,
        },
        success: function (res) {
            if (res.status === 200) {

                res.data.forEach(
                    function (val) {
                        //console.log(val);
                        // 看起来是深拷贝, 借助引用地址的样子
                        val.headimg = myApi + "/" + val.headimg;
                    }
                );

                // 加载聊天内容
                // 设置聊天内容样式
                setMsgContent(res.data);

            }
        }
    })
}

// 设置聊天内容样式
function setMsgContent(data) {

    if (data instanceof Array) {
        // 数组
        data.forEach(
            function (val) {
                // 判断是否为当前用户
                // 获取个人信息
                const myInfo = JSON.parse(Cookies.get("myInfo"));
                // 判断是否为当前用户
                if (myInfo.id === val.from_uid) {
                    $('<li class="sent"><img src="' + val.headimg + '" alt="" /><p>' + val.content + '</p></li>').appendTo($('.messages ul'));
                } else {
                    $('<li class="replies"><img src="' + val.headimg + '" alt="" /><p>' + val.content + '</p></li>').appendTo($('.messages ul'));
                }

                $('.message-input input').val(null);
                $('.contact.active .preview').html('<span>You: </span>' + val.content);
                $(".messages").animate({scrollTop: $(document).height()}, "fast");
            }
        );
        return
    }

    // 判断是否为当前用户
    // 获取个人信息
    const myInfo = JSON.parse(Cookies.get("myInfo"));
    // 判断是否为当前用户
    if (myInfo.id === data.from_uid) {
        $('<li class="sent"><img src="' + data.headimg + '" alt="" /><p>' + data.content + '</p></li>').appendTo($('.messages ul'));
    } else {
        $('<li class="replies"><img src="' + data.headimg + '" alt="" /><p>' + data.content + '</p></li>').appendTo($('.messages ul'));
    }

    $('.message-input input').val(null);
    $('.contact.active .preview').html('<span>You: </span>' + data.content);
    $(".messages").animate({scrollTop: $(document).height()}, "fast");

}

// 绑定聊天事件
$('.submit').click(function () {
    newMessage();
});

$(window).on('keydown', function (e) {
    if (e.which === 13) {
        newMessage();
        return false;
    }
});

// 组成消息体
// 发送聊天消息
function newMessage() {

    // 获取个人信息
    const myInfo = JSON.parse(Cookies.get("myInfo"));

    // 群聊id
    groupId = $("#groupUser").find(".groupId").text();

    // 发送消息体
    const content = $(".message-input input").val();
    if ($.trim(content) === '') {
        return false;
    }

    msgData = {
        name: myInfo.name,
        headimg: myInfo.headimg,
        content: content,
        //测试
        group_id: groupId,
        from_uid: myInfo.id,
        content_type: "text"

    };

    // ws消息发送
    // 发送数据
    console.log(msgData);
    // 发送体必须为string
    // json会出错
    ws.send(JSON.stringify(msgData));
}


// ===========================================
// 添加好友
// ===========================================

// 用户搜索
$("#searchFriend").on("keydown", function (e) {
    //alert($("#searchFriend input").val());
    if (e.which === 13) {
        // 获取内容
        const content = $("#searchFriend input").val();
        if ($.trim(content) === '') {
            return false;
        }

        // 请求数据
        // 用户id
        const uid = Cookies.get("uid")

        // 请求数据
        $.ajax({
            url: myApi + "/user/search",
            method: "GET",
            data: {
                every: "all",
                // 关键字搜索
                key: content,
            },
            success: function (res) {
                if (res.status === 200) {
                    // 修改数据
                    // 通过模板引擎渲染数据

                    res.data.forEach(
                        function (val) {
                            val.headimg = myApi + "/" + val.headimg;
                        }
                    );

                    if (res.data.length === 0) {
                        // 清除数据
                        $("#contact").empty();
                    }

                    // 渲染数据
                    $("#friends").html($("#addUserList").render(res.data))

                } else {
                    confirm(res.msg)
                }
            }
        })
    }
});

// 添加好友
// 群聊创建
function disGroup(friendId) {

    var uids = friendId;
    uids += "," + Cookies.get("uid");
    //alert(uids);
    // 请求数据
    $.ajax({
        url: myApi + "/chat/disGroup",
        method: "POST",
        data: {
            uids: uids,
        },
        success: function (res) {
            if (res.status === 200) {
                // 修改数据
                // 通过模板引擎渲染数据

            } else {
                confirm(res.msg)
            }
        }
    })
}