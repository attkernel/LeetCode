package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
)

var Base36 = "0123456789abcdefghijklmnopqrstuvwxyz"

func main() {
	fileInfo, _ := os.Stdin.Stat()
	if (fileInfo.Mode() & os.ModeNamedPipe) != os.ModeNamedPipe {
		log.Fatal("The command is intended to work with pipes.")
	}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		fmt.Println(Convert36(s.Text()))
	}
	/*fmt.Println(Convert36("YQ=="))
	fmt.Println(Convert36("eW8="))
	fmt.Println(Convert36("cGFzc3dvcmQgaXMgQWRtaW4xMjM0NTY="))*/
}

func NumTo36Hex(num *big.Int) string {
	str := ""
	//循环取余,根据余值对应base36字符数组内容
	for num.Cmp(big.NewInt(0)) != 0 {
		tmp := big.NewInt(0)
		tmp.Mod(num, big.NewInt(36))
		str = string(Base36[tmp.Int64()]) + str
		num.Div(num, big.NewInt(36))
	}
	return strings.ToUpper(str)
}

func Convert36(src string) string {
	hexStr := ""
	//base64解密
	data, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return ""
	}
	//将解密的原文拼接成16进制字符串
	for idx, _ := range data {
		str := strconv.FormatInt(int64(data[idx]), 16)
		hexStr += str
	}
	//防止溢出,使用bigInt,将16进制字符串读取成bigInt
	num := new(big.Int)
	num.SetString(hexStr, 16)
	return NumTo36Hex(num)
}
