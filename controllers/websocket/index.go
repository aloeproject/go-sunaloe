package webim

import (
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"net/http"
	"myweb/service/websocket"
	"fmt"
)

type IndexController struct {
	beego.Controller
}


func (this *IndexController) WSocket(){

	ws,err := websocket.Upgrade(this.Ctx.ResponseWriter,this.Ctx.Request,nil,1024,1024)
	if _,ok := err.(websocket.HandshakeError); ok {
		http.Error(this.Ctx.ResponseWriter,"不是一个websocket 连接",400)
		return
	} else if err != nil {
		http.Error(this.Ctx.ResponseWriter,"websocket 连接失败",400)
		return
	}
	userHash,_ := this.Ctx.GetSecureCookie(beego.AppConfig.String("cookie.secure"),"uuid")
	//进入聊天室
	//先不做登录用户 都是匿名的
	userId := 0
	userName := service.GetAnonymousName(userHash)
	anonymousId := service.GetAnonymousId(userHash)
	webimUserSession := fmt.Sprintf("webim-%s",userHash)
	var userObj  service.SocketUser
	if webSession:=this.GetSession(webimUserSession);webSession == nil {
		userObj = service.SocketUser{User_hash:userHash,User_id:userId,User_name:userName,Is_login:false,AnonymousId:anonymousId}
		this.SetSession(webimUserSession,userObj)
	} else {
		//interface{} 转换为 SocketUser类型
		switch v := webSession.(type) {
		case service.SocketUser:
			userObj = v
			userObj.Is_login = true
		}
	}

	fmt.Println(userObj,userHash)
	//进入聊天室
	service.JoinRoom(userObj,ws)
	//离开聊天室
	defer service.LeaveRoom(userObj)


	for {
		_,p,err := ws.ReadMessage()
		if err != nil {
			return
		}
		service.Publish <- service.NewWsEvent(service.EVENT_MESSAGE,userObj,string(p))
	}
}