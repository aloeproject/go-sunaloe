package service

import (
	"github.com/gorilla/websocket"
	"container/list"
	"encoding/json"
	"github.com/astaxie/beego"
	"fmt"
	"log"
)
var thisUserHash = ""
//用户结构
type SocketUser struct{
	User_name string
	User_hash string
	User_id int
	Is_login bool
	AnonymousId int
}

//订阅者结构
type Subscriber struct {
	User SocketUser
	Conn *websocket.Conn
}

//进入聊天室
func JoinRoom(user SocketUser,ws *websocket.Conn,userhash string){
	thisUserHash = userhash
	subscribe <- Subscriber{User:user,Conn:ws}
}
//离开聊天室
func LeaveRoom(user SocketUser) {
	unsubscribe <- user
}

func GetAnonymousId(userHash string) int {
	//从最后 往前找用户
	for sub:=subscribers.Front();sub!=nil;sub = sub.Next() {
		if sub.Value.(Subscriber).User.User_hash == userHash {
			uid := sub.Value.(Subscriber).User.User_id
			return uid
		}
	}

	for sub:=subscribers.Back();sub!=nil;sub = sub.Prev() {
		id := sub.Value.(Subscriber).User.AnonymousId
		return id+1
	}
	return 1
}

//得到游客姓名
func GetAnonymousName(userHash string) string{
	for sub:=subscribers.Front();sub != nil;sub = sub.Next() {
		if sub.Value.(Subscriber).User.User_hash == userHash {
			name := sub.Value.(Subscriber).User.User_name
			return name
		}
	}
	//从最后 往前找用户
	for sub:=subscribers.Back();sub!=nil;sub = sub.Prev() {
		id := sub.Value.(Subscriber).User.AnonymousId
		return fmt.Sprintf("游客:%d",id + 1)
	}
	return fmt.Sprintf("游客:%d",1)
}


var (
	//同时10个登录
	subscribe = make (chan Subscriber,10)
	//同时10个退出
	unsubscribe = make(chan SocketUser,10)
	//publish
	Publish = make(chan wsEvent,10)
	//等待和订阅列表
	subscribers = list.New()

)



func wsRoom(){
	for {
		select {
		//有新的用户加入
		case sub := <- subscribe:
			//fmt.Println(subscribers.Len())
			//这里进行用户判断重复，或者老用户登录
			if isUserExist(subscribers,sub) {
				//info := fmt.Sprintf("%s 欢迎再次回来",sub.User.User_name)
				//Publish <- NewWsEvent(EVENT_MESSAGE,sub.User,info)
				log.Println("老用户")
			} else {
				subscribers.PushBack(sub)
				length := subscribers.Len()
				info := ""
				if sub.User.Is_login == false {
					info = fmt.Sprintf("%s 欢迎进入聊天室，当前人数 %d 人",sub.User.User_name,length)
					Publish <- NewWsEvent(EVENT_JOIN,sub.User,info)
				} else {
					info = fmt.Sprintf("%s 上线了，当前人数 %d 人",sub.User.User_name,length)
					Publish <- NewWsEvent(EVENT_JOIN,sub.User,info)
				}
			}
		case event := <- Publish:
			//广播通知
			broadcastWebSocket(event)

		case unUser := <-unsubscribe:
			for sub:= subscribers.Front();sub != nil;sub = sub.Next() {
				//如果某个用户退出 hash 来标识用户唯一
				if sub.Value.(Subscriber).User.User_hash == unUser.User_hash {
					//列表剔除改用户
					subscribers.Remove(sub)
					ws := sub.Value.(Subscriber).Conn
					if ws != nil {
						//关闭此用户连接
						ws.Close()
					}
					msg := fmt.Sprintf("用户:%s 下线了",unUser.User_name)
					Publish <- NewWsEvent(EVENT_LEAVE,unUser,msg)
					break
				}

			}
		}

	}
}

func init()  {
	go wsRoom()
}
//广播
func broadcastWebSocket(event wsEvent){

	//循环通知订阅者
	for sub := subscribers.Front();sub != nil;sub = sub.Next() {
		ws := sub.Value.(Subscriber).Conn
		if sub.Value.(Subscriber).User.User_hash == thisUserHash {
			event.User.User_name = fmt.Sprintf("我(%s)",event.User.User_name)
		}
		data,err := json.Marshal(event)
		if err != nil {
			beego.Error("fail to marshal event:",err)
			return
		}
		if ws != nil {
			//广播通知，如果通知失败则已经离开房间
			if ws.WriteMessage(websocket.TextMessage,data) != nil {

			}
		}
	}
}
//判断用户是否存在
func isUserExist(subscribers *list.List,user Subscriber) bool{
	for sub:=subscribers.Front();sub != nil;sub = sub.Next() {
		if sub.Value.(Subscriber).User.User_hash == user.User.User_hash {
			return true
		}
	}
	return false
}