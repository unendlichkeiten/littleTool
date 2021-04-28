package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var (
	length  int
	charSet string
)

const (
	NumStr  = "0123456789"
	CharStr = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	SpecStr = "+=-@#~,.[]()!%^*$"
)

func main() {
	// 解析入参
	flag.IntVar(&length, "n", 8, "-l 密钥长度")
	flag.StringVar(&charSet, "t", "num",
		`-t 生成密码的字符集
num  : 数字密码
char : 字符密码
mix  : 混合密码（包含数字、普通字符和特殊字符）`)

	flag.Parse()

	// 生成创建密钥的原始串
	var srcStr string
	var passwd []byte = make([]byte, length, length)
	switch charSet {
	case "num":
		srcStr = NumStr
	case "char":
		srcStr = CharStr
	case "mix":
		srcStr = NumStr + CharStr + SpecStr
	default:
		srcStr = NumStr
	}

	// 生成密钥
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(passwd); i++ {
		passwd[i] = srcStr[rand.Intn(len(srcStr))]
	}

	fmt.Printf("passwd: %s", string(passwd))
}
