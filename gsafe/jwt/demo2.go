package main

import (
	"io/ioutil"
	"fmt"
	"crypto/sha256"
	"crypto/md5"
	"hash"
	"encoding/hex"
)

func Sha256(src []byte) string {
	return encode(sha256.New(), src)
}

func Md5(src []byte) string {
	return encode(md5.New(), src)
}

func encode(h hash.Hash, src []byte) string {
	h.Write(src)
	return hex.EncodeToString(h.Sum(nil))
}

func main() {
	file := "1.3.2.uns"
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}
	shar := Sha256(data)
	fmt.Println(shar)
}
