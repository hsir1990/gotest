package utils

import (
	"crypto/md5"
	"fmt"
)

func Md5(unEncry string) (encryption string) {
	byte16 := []byte(unEncry)
	encryption = fmt.Sprintf("%x", md5.Sum(byte16)) //%x	表示为十六进制，使用a-f
	return
}

// func Sum
// func Sum(data []byte) [Size]byte
// 返回数据data的MD5校验和。
