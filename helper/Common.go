package helper

import (
	"fmt"
	"strconv"
	"time"
	"crypto/md5"
	"encoding/hex"
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


