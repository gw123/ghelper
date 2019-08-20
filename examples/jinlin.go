package main

import (
	"crypto/md5"
	"io"
	"fmt"
	"flag"
)

func generateNewPassword(productKey string, productName string) (password string) {
	h := md5.New()
	io.WriteString(h, productKey)
	io.WriteString(h, productName)
	password = fmt.Sprintf("%x", h.Sum(nil))
	return
}

func main() {
	pname := flag.String("name", "", "设备名称")
	flag.Parse()
	if *pname == "" {
		fmt.Println("设备名称不能为空")
		return
	}
	pwd := generateNewPassword("a1deuUtiPdt", *pname)
	fmt.Println(pwd)
}
