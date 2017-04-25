package service

import (
	"github.com/gorilla/websocket"
	"time"
	"myweb/helper"
	"encoding/json"
	"log"
	"fmt"
)


var (
	//同时10个登录
	connecter = make(chan Connecter, 10)
	unconnect = make(chan Connecter, 10)
	DataPushlish = make(chan Connecter,10)
	timeAnddate = make(map[int64]int)

	//数据的时间跨度
	timespan = NowUnix(60)
	totalTime = NowUnix(3600)
)

type NowUnix int

type Connecter struct {
	nowtime NowUnix
	Conn *websocket.Conn
}


type Webrealtimedata struct {
	Data_time []string
	Real_count []int
}

func NewConnecter(ws *websocket.Conn) Connecter {
	n := NowUnix(time.Now().Unix())
	return Connecter{n,ws}
}

//加入连接
func JoinConn(ws *websocket.Conn){
	connecter <- NewConnecter(ws)
}
//断开连接
func CloseConn(ws *websocket.Conn) {
	unconnect <- NewConnecter(ws)
}

func init(){
	go clearTimeAnddate()
	go realWork()
}


func realWork(){
	for {
		select {
			//连接
			case  <- connecter:
				t := time.Now().Unix()
				timeAnddate[t] += 1

			//send
			case	pu := <- DataPushlish:
				var webdata Webrealtimedata
				//总时间
				for i := NowUnix(0);i<=totalTime;i+= timespan {
					t := int64(pu.nowtime - i)
					dataTime := helper.GetDate(t,"15:04:05")
					realcount := getRealTotalCount(int64(pu.nowtime - totalTime),int64(pu.nowtime - i))
					webdata.Data_time = append(webdata.Data_time,dataTime)
					webdata.Real_count = append(webdata.Real_count,realcount)
				}

				fmt.Println(timeAnddate)

				data,err := json.Marshal(webdata)
				if err != nil {
					log.Println("fail to marshal",err)
				}

				ws := pu.Conn
				if ws.WriteMessage(websocket.TextMessage,data) != nil {
					log.Println("send message fail",err)
				}

			//端口
			case  un:=<- unconnect:
				fmt.Println("--------")
				t := time.Now().Unix()
				timeAnddate[t] -= 1
				ws := un.Conn
				if ws != nil {
					//关闭此用户连接
					ws.Close()
				}
		}

	}
}
//避免timeAnddate 越来越大
func clearTimeAnddate(){
	//一小时清理一次
	ed := time.Now().Unix()
	st := int64(NowUnix(ed) - totalTime)
	tmpTimeAnddate := make(map[int64]int)
	for i:= st;i<=ed;i++ {
		if timeAnddate[i] != 0 {
			tmpTimeAnddate[i] = timeAnddate[i]
		}
	}
	timeAnddate = tmpTimeAnddate
}


func getRealTotalCount(startTime int64,endTime int64) (int){
	//fmt.Println(timeAnddate)
	count := 0
	for i:=startTime;i<=endTime;i++ {
		key := int64(i)
		if timeAnddate[key] != 0 {
			count = timeAnddate[key] + count
		}
	}
	return count
}
