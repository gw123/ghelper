package reg_util

func main() {
	//nums := []string{"ss0.01","10.","100&","10.1", "0.02", "12132.00", "1000000.02"}
	//for _, num := range nums{
	//	reg := regexp.MustCompile(`(\d+(\.\d+)?)`)
	//	if reg.MatchString(num) {
	//		fmt.Println("match")
	//	}
	//}

	//匹配 201901001  或者 201910 这样的日期 必须是 8位或者6位
	//reg := regexp.MustCompile(`^(\d{6}|\d{8})$`)
	//if reg.MatchString("201809") {
	//	fmt.Println("match")
	//} else {
	//	fmt.Println("not match")
	//}


	//b := "rpc error: grpc_code = InvalidArgument desc = 未找到设备 请重试"
	//reg := regexp.MustCompile(`rpc error: grpc_code = ([\w\d]+) desc = (.*)`)
	////res := reg.FindAllStringSubmatch(b,-1)
	//res := reg.FindStringSubmatch(b)
	//glog.Dump(res)

}
