package helper

import "time"

func GetNowDate() string  {
	loc, _:= time.LoadLocation("Asia/Chongqing")
	t := time.Now().In(loc)
	return t.Format("2006-01-02 15:04:05")
}
/*
func Mktime(hour,minute,second,month,day,year int){
	loc, _:= time.LoadLocation("Asia/Chongqing")
	layout := "2006-01-02 15:04:05"

}
*/