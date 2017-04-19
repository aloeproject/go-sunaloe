package vendor

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
)

const salt  = "$J)JIK)D#WD#!@@E@#a"


func GetPasswordHash(str string) string {
	md5Inst := md5.New()
	md5Inst.Write([]byte(str))
	md5Ret := hex.EncodeToString(md5Inst.Sum([]byte(salt)))


	sha1Inst := sha1.New()
	sha1Inst.Write([]byte(md5Ret))
	passwordHash := hex.EncodeToString(sha1Inst.Sum(nil))
	return string(passwordHash)
}

func CheckPassword(password,password_hash string) bool{
	checkPd := GetPasswordHash(password)
	return password_hash == checkPd
}
