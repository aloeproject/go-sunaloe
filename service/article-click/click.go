package article_click

import (
	"github.com/astaxie/beego/cache"
	"time"
	"myweb/repository"
	"log"
	"strconv"
	"fmt"
)

var (
	mycache cache.Cache
)

func init() {
	mycache, _ = cache.NewCache("memory", `{"interval":0}`)
}


func  SetClick(aid int,uuid string,ip string)  {
	k := fmt.Sprintf("article_click:%d-%s",aid,uuid)
	if mycache.IsExist(k) == false {
		//插入的点击半小时算重新阅读
		mycache.Put(k,1,1800 * time.Second)
		rep := new(repository.ArticleClickRepository)
		_,err := rep.Save(aid,ip,uuid)
		if err != nil {
			log.Print("插入点击量失败",err)
		}
	}
}

func  GetClick(aid int) int {
	rep := repository.ArticleClickRepository{aid}
	k := fmt.Sprintf("article_click_num:%d",aid)
	if mycache.IsExist(k) == true {
		numStr := mycache.Get(k)
		num,err := strconv.Atoi(fmt.Sprint(numStr))
		if err != nil {
			log.Print("缓存获得点击数失败")
		} else {
			return num
		}
	}
	num,err := rep.GetClickCount()
	if err != nil {
		log.Print("获得点击率失败")
		return 0
	} else {
		//点击率30秒更新一次
		if err := mycache.Put(k,num,30*time.Second);err != nil {
			log.Print("设置点击数失败aid:",aid)
		}
		return num
	}
}
