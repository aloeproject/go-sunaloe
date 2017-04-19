package library

import (
	"math"
	"fmt"
	"strconv"
)

//分页

type Page struct {
	PageNo int //当前页
	PageSize int
	TotalPage int
	TotalCount int //总记录数
	FirstPage bool
	LastPage bool
	List interface{}
}

func NewPage(count int,pageNo int,pageSize int,list interface{}) Page {
	tp := math.Ceil(float64(count) / float64(pageSize)) //总页数
	totalPage,_ := strconv.Atoi(fmt.Sprint(tp))
	return Page{PageNo:pageNo,PageSize:pageSize,TotalPage:totalPage,TotalCount:count,
		FirstPage:pageNo == 1,LastPage:pageNo >= totalPage,List:list}
}
