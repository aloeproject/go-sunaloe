package library

import (
	"fmt"
	"strconv"
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

