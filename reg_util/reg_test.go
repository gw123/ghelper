package reg_util

import "testing"

func TestIsIP2(t *testing.T) {
	ips := []string{"1.1.1.1", "0.0.0.0", "112.112.112.112", "255.255.1.1"}

	for _, ip := range ips {
		if IsIP(ip) != true {
			t.Fail()
		}
	}

	notIps := []string{"1.1.1.1.1", "0.0.0.", "112.112.112.1122", "2552.255.1.1",
		"90.1112.22.2", "90.1112.22.21", "1a.111.22.2",
	}

	for _, ip := range notIps {
		if IsIP(ip) == true {
			t.Fail()
		}
	}
	t.Log("IsIP2 pass")
}

func TestIsVersionStr(t *testing.T) {
	ips := []string{"1.1.1", "0.0.0", "112.112.112", "1.2.15"}

	for _, ip := range ips {
		if IsVersionStr(ip) != true {
			t.Fail()
		}
	}

	notIps := []string{"a.1.1.1.1", "0.0.0.", "112.112.c", "2552.255.1.1",
		"90.1112.22.2", "90.1112.22.21", "1a.111.22.2",
	}

	for _, ip := range notIps {
		if IsVersionStr(ip) == true {
			t.Fail()
		}
	}
	t.Log("IsVersionStr pass")
}

func TestGetEvnVar(t *testing.T) {
	res, ok := GetEvnVar("${DB_NAME}")
	if !ok {
		t.Fail()
	}
	t.Log("MATCH sub : ", res)
}

func TestFindStringFirstNum(t *testing.T) {
	testStrs := []map[string]string{
		{
			"input": "10.2",
			"res":   "10.2",
		},
		{
			"input": "收款10元",
			"res":   "10",
		},
		{
			"input": "卖了101个",
			"res":   "101",
		},
	}

	for _, obj := range testStrs {
		num := FindStringFirstNum(obj["input"])
		if obj["res"] != num {
			t.Errorf("解析失败 num:%s res:%s", num, obj["res"])
		}
	}
}

func TestIsDateNumber(t *testing.T) {
	testStrs := []string{
		"201908",
		"20191111",
		"20181212",
	}
	for _, str := range testStrs {
		ok := IsDateNumber(str)
		if !ok {
			t.Fatal("IsDateNumber() 转换失败")
		}
		t.Logf("转换成功")
	}
}

func TestParseGrpcError(t *testing.T) {
	testStrs := []string{
		//"rpc error: grpc_code = InvalidArgument desc = 未找到设备 请重试",
		//"rpc error: grpc_code = DeadlineExceeded desc = 请求超时",
		"rpc error: grpc_code = AlreadyExists desc = 资源已经存在",
		"rpc error: code = Code(1026) desc = 未找到设备",
	}
	for _, str := range testStrs {
		code, desc, ok := ParseGrpcError(str)
		if !ok {
			t.Fatal("ParseGrpcError() 转换失败")
		}
		t.Logf("grpc_code : %s , desc : %s", code, desc)
	}

	testStrs2 := []string{
		//"rpc error: code1 = InvalidArgument desc = 未找到设备 请重试",
		//"rpc error: grpc_code = DeadlineExceeded ",
		//"rpc error: grpc_code desc = 资源已经存在",
		"rpc error: code = Code123(1026) desc = 未找到设备",
	}
	for _, str := range testStrs2 {
		_, _, ok := ParseGrpcError(str)
		if ok {
			t.Fatal("ParseGrpcError() 解析错误")
		}
	}
}
