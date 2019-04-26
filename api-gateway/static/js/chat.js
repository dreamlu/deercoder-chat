//var myObj = {}
// var
var myApi = "http://localhost:8006/api/v1"

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
                // 渲染数据
                var tmpl = $.template(".contact")
                tmpl.render(eval(res.data))
                $("#contacts").html($("#userList").render(eval(res.data)))

            } else {
                confirm(res.msg)
            }
        }
    })
}