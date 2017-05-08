package helper

import (
	"fmt"
	"strconv"
	"time"
	"crypto/md5"
	"encoding/hex"
	"os"
)

/**
   字符串转整型
 */
func String2int(num string) (int,error) {
	ct := fmt.Sprint(num)
	i ,err := strconv.Atoi(ct)
	if err == nil {
		return i,nil
	}
	return 0,err
}
/**
 得到唯一uuid
 */
func GetUUID() (string){
	unixNano :=fmt.Sprint(time.Now().UnixNano())
	md5Inst := md5.New()
	md5Inst.Write([]byte(unixNano))
	md5Ret := hex.EncodeToString(md5Inst.Sum([]byte("")))

	return fmt.Sprintf("%s-%s-%s-%s-%s",md5Ret[0:8],md5Ret[8:12],md5Ret[12:16],md5Ret[16:20],md5Ret[20:])
}
/**
  判断目录是否存在
 */
func PathExists(path string) (bool,error)  {
	_,err := os.Stat(path)
	if err == nil {
		return true,nil
	}
	if os.IsNotExist(err) {
		return false,nil
	}
	return false,err
}

func GetUploadImageDir() (string,error) {
	uploadDir := "./static/upload/image"
	t := time.Now()
	//年和月新建文件夹
	uploadDir = uploadDir+"/"+t.Format("200601")
	isExistDir,err := PathExists(uploadDir)
	//目录不存在则创建
	if isExistDir == false {
		err = os.Mkdir(uploadDir,os.ModePerm)
	}
	if err != nil {
		return "",err
	}
	return uploadDir,nil
}

