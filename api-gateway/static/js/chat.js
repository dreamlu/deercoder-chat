
// 登陆
function login() {
    var name = $("#name").val()
    var password = $("#password").val()

    $.ajax({
        url: "/api/v1/login/login",
        method: "POST",
        data:{
            name: name,
            password: password
        },
        success: function (res) {
            console.log(res)
        }
    })

}