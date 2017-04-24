$(function () {

    var webIm = {
        "socket":null,
        "init":function () {
            // Create a socket
            //socket = new WebSocket('ws://' + 'localhost:8080' + '/ws/join?uname=123');
            socket = new WebSocket('ws://' + 'localhost:8081' + '/im');
            this.socket = socket
            // Message received on the socket
            socket.onmessage = function (event) {
                var data = JSON.parse(event.data);
                $("#im-content-box").append('<div class="from-message" data-msg-id="'+data.Timestamp+'">'+data.User.User_name+' ：<div class="wrapper">'+data.Content+'</div></div>')
            };

        },
        "send":function (msg) {
            this.socket.send(msg)
        }
    }

    $("#anonymous-im").on("click",function () {
        $(".im-chat-window").show()
    })
    $("body").on("click",function () {
        $(this)
        //$(".im-chat-window").hide()
    })
    
    $("#im-send-msg").on("click",function () {
        var content = $("#im-content").val()
        webIm.send(content)
        $("#im-content").val("")
        //滚动条到底部
        var height = $("#im-content-box")[0].scrollHeight
        $("#im-content-box").scrollTop(height * 1.1)
    })

    webIm.init()
})
