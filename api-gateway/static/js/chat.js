
// 登陆
function login() {
    var name = $("#name").val();
    var password = $("#password").val();

    $.ajax({
        url: "/api/v1/login/login",
        method: "POST",
        data:{
            name: name,
            password: password
        },
        success: function (res) {
            console.log(res)
            if(res.status === 200) {
                location.href = "index.html"
            }
            else if(res.status === 211){
                confirm(res.msg)
            }
        }
    })

}