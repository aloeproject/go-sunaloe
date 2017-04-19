package library

import "time"

func GetNowDate() string  {
	loc, _:= time.LoadLocation("Asia/Chongqing")
	t := time.Now().In(loc)
	return t.Format("2006-01-02 15:04:05")
}
