$(function () {
    var webIm = {
        "socket":null,
        "init":function () {
            // Create a socket
            //socket = new WebSocket('ws://' + 'localhost:8080' + '/ws/join?uname=123');
            socket = new WebSocket('ws://' + window.location.host  + '/web_data');
            this.socket = socket
            // Message received on the socket
            socket.onmessage = function (event) {
                var data = JSON.parse(event.data);
                console.log(data)
            };

        },
    }
    webIm.init()
})

