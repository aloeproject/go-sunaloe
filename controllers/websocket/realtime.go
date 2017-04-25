package webim

import (
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"net/http"
	//"time"
	"myweb/service/websocket"
	"time"
	"fmt"
)

type RealtimeController struct {
	beego.Controller
}

func (this *RealtimeController)SetOnlineData()  {
	ws,err := websocket.Upgrade(this.Ctx.ResponseWriter,this.Ctx.Request,nil,1024,1024)
	if _,ok := err.(websocket.HandshakeError);ok {
		http.Error(this.Ctx.ResponseWriter,"不是一个websocket连接",400)
		return
	} else if err != nil {
		http.Error(this.Ctx.ResponseWriter,"websocket 连接失败",400)
		return
	}

	defer service.CloseConn(ws)

	ws.ReadMessage()
	fmt.Println(1111)
	service.JoinConn(ws)

}

func (this *RealtimeController) GetOnlineData()  {
	ws,err := websocket.Upgrade(this.Ctx.ResponseWriter,this.Ctx.Request,nil,1024,1024)
	if _,ok := err.(websocket.HandshakeError);ok {
		http.Error(this.Ctx.ResponseWriter,"不是一个websocket连接",400)
		return
	} else if err != nil {
		http.Error(this.Ctx.ResponseWriter,"websocket 连接失败",400)
		return
	}


	for {
		service.DataPushlish <- service.NewConnecter(ws)
		time.Sleep(2 * time.Second)
	}

}