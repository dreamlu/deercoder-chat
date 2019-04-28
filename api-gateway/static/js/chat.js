//var myObj = {}
// var
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

    // 好友信息
    // 聊天界面用户头像名称等信息
    const arrayData = {
        'name': name,
        'headimg': headimg
    };
    $("#groupUser").html($("#groupUserInfo").render(arrayData));

    // 建立在线websocket即时通讯
    // 开始聊天
    if ("WebSocket" in window) {
        ws = new WebSocket(myWsApi + "/chat/chatWs");

        ws.onopen = function () {
            // 拉取离线数据

            // // 发送数据
            // ws.send(JSON.stringify({
            //         name: name,
            //         headimg: headimg,
            //         content: "聊天测试",
            //         //测试
            //         group_id: "93f65451-efc4-11e8-918b-34e6d7558045",
            //         from_uid: 1,
            //         content_type: "text"
            //     }
            // ));
            //alert("数据发送中...");
        };

        // 接收数据
        ws.onmessage = function (res) {
            console.log("[receive data]: " + res.data);
            const data = JSON.parse(res.data);
            //data.headimg = myApi + "/" + data.headimg;
            //alert("数据已接收...");
            $('<li class="sent"><img src="' + data.headimg + '" alt="" /><p>' + data.content + '</p></li>').appendTo($('.messages ul'));
            $('.message-input input').val(null);
            $('.contact.active .preview').html('<span>You: </span>' + data.content);
            $(".messages").animate({scrollTop: $(document).height()}, "fast");
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


// 绑定聊天事件
$('.submit').click(function () {
    newMessage();
});

$(window).on('keydown', function (e) {
    if (e.which == 13) {
        newMessage();
        return false;
    }
});

// 组成消息体
function newMessage() {

    // 获取个人信息
    const myInfo = JSON.parse(Cookies.get("myInfo"));

    // 发送消息体
    const content = $(".message-input input").val();
    msgData = {
        name: myInfo.name,
        headimg: myInfo.headimg,
        content: content,
        //测试
        group_id: "93f65451-efc4-11e8-918b-34e6d7558045",
        from_uid: myInfo.id,
        content_type: "text"

    };
    if ($.trim(content) === '') {
        return false;
    }

    // ws消息发送
    // 发送数据
    console.log(msgData);
    // 发送体必须为string
    // json会出错
    ws.send(JSON.stringify(msgData));
}