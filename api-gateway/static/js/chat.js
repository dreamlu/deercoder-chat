// api地址等信息
var api = "localhost:8006";
var httpProtocol = "http";
var wsProtocol = "ws";
var myApi = httpProtocol + "://" + api + "/api/v1";
var myWsApi = wsProtocol + "://" + api + "/api/v1";

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

                // 绑定点击个人数据展示修改页面

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

    newWebSocket();

});

// 建立websocket
function newWebSocket(){
    // 建立在线websocket即时通讯
    // 开始聊天
    if ("WebSocket" in window) {
        ws = new WebSocket(myWsApi + "/chat/chatWs");

        // 建立ws连接
        ws.onopen = function () {

            // 关闭之前的连接
            //ws.close();

            console.log("连接建立...");
            heartCheck.reset().start();   // 成功建立连接后，重置心跳检测
            // 拉取离线数据

            // 绑定文件上传事件
            setUpLoadFile();
        };

        // 接收数据
        ws.onmessage = function (res) {

            heartCheck.reset().start(); // 如果获取到消息，说明连接是正常的，重置心跳检测
            // console.log("[receive data]: " + res.data);
            // 返回的数据json转换
            const data = JSON.parse(res.data);
            // 设置聊天内容样式
            setMsgContent(data);
        };

        // 连接关闭
        ws.onclose = function () {
            // 关闭 websocket
            console.log("连接已关闭...");
        };
    } else {
        // 浏览器不支持 WebSocket
        alert("您的浏览器不支持 WebSocket!");
    }
}

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
                // 聊天内容样式解析
                // 文件类型
                let fileType;
                if (val.content_type === 'file') {
                    // 文件具体类型
                    fileType = val.content.split(".")[1];
                    // 图片类型
                    if (fileType === ("png" || "jpg" || "jpeg")) {
                        val.content = '<img style="width: 200px; border-radius: 0;" src="' + myApi + "/" + val.content + '" alt="文件内容" />';
                    }
                    // 其他文件下载类型
                    // 待完善, 有兴趣的可以完善
                    // 显示文件名, 下载按钮
                    // 可利用content字段组合文件链接
                    // 解析显示下载文件
                }

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

    // 聊天内容样式解析
    // 文件类型
    let fileType;
    if (data.content_type === 'file') {
        // 文件具体类型
        fileType = data.content.split(".")[1];
        // 图片类型
        if (fileType === ("png" || "jpg" || "jpeg")) {
            data.content = '<img style="width: 200px; border-radius: 0;" src="' + myApi + "/" + data.content + '" alt="文件内容" />';
        }
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

// 文件类型
// 组成消息体
// 发送聊天消息
function newMessageFile(filePath) {

    // 获取个人信息
    const myInfo = JSON.parse(Cookies.get("myInfo"));

    // 群聊id
    groupId = $("#groupUser").find(".groupId").text();

    // 发送消息体
    msgData = {
        name: myInfo.name,
        headimg: myInfo.headimg,
        content: filePath,
        //测试
        group_id: groupId,
        from_uid: myInfo.id,
        content_type: "file"

    };

    // ws消息发送
    // 发送数据
    console.log(msgData);
    // 发送体必须为string
    // json会出错
    ws.send(JSON.stringify(msgData));
}

// 聊天内容
// 文件类型
function setUpLoadFile(){
    // 聊天内容
    // 文件类型
    $("#fileUp").after('<input type="file" id="upLoadFile" name="file" style="display:none" onchange ="uploadFile()">');
    $("#fileUp").click(function () {
        // 点击上传
        document.getElementById("upLoadFile").click();
    });
}

// 文件上传
function uploadFile() {
    var myform = new FormData();
    myform.append('file', $('#upLoadFile')[0].files[0]);
    $.ajax({
        url: myApi + "/file/upload",
        type: "POST",
        data: myform,
        contentType: false,
        processData: false,
        success: function (res) {
            //console.log(res);
            // 文件上传成功
            if(res.status === 224){
                //return res.path;
                // 发送文件类型消息
                newMessageFile(res.path);
            }
        }
    });
    return null;
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
                if (res.data === "好友已存在") {
                    confirm(res.data)
                }
            }
        }
    })
}


// 心跳检测, 每隔一段时间检测连接状态，如果处于连接中，就向server端主动发送消息，来重置server端与客户端的最大连接时间，如果已经断开了，发起重连。
var heartCheck = {
    timeout: 1000 * 60 * 9,        // 9分钟发一次心跳，比server端设置的连接时间稍微小一点，在接近断开的情况下以通信的方式去重置连接时间。
    serverTimeoutObj: null,
    reset: function(){
        clearTimeout(this.timeoutObj);
        clearTimeout(this.serverTimeoutObj);
        return this;
    },
    start: function(){
        var self = this;
        this.serverTimeoutObj = setInterval(function(){
            if(ws.readyState === 1){
                console.log("连接状态，发送消息保持连接");
                ws.send("ping");
                heartCheck.reset().start();    // 如果获取到消息，说明连接是正常的，重置心跳检测
            }else{
                console.log("断开状态，尝试重连");
                newWebSocket();
            }
        }, this.timeout)
    }
};
