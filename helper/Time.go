package helper

import "time"

func GetNowDate() string  {
	loc, _:= time.LoadLocation("Asia/Chongqing")
	t := time.Now().In(loc)
	return t.Format("2006-01-02 15:04:05")
}

func GetDate(unixtime int64,format string) string  {
	tm := time.Unix(unixtime,0)
	loc, _:= time.LoadLocation("Asia/Chongqing")
	return tm.In(loc).Format(format)
}