package main

import (
	"fmt"
	"github.com/gw123/glog"
	"runtime"
)

func main() {
	data := make([]int, 1024)
	stat := &runtime.MemStats{}
	runtime.ReadMemStats(stat)

	glog.Dump(stat)

	fmt.Println(data[0])
}
