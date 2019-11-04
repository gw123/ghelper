package reg_util

import (
	"regexp"
	"strings"
)

/***
  ip 格式[10.0.0.1]
*/
func IsIP(input string) bool {
	exp1 := regexp.MustCompile(`^\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}$`)
	return exp1.MatchString(input)
}

func IsVersionStr(input string) bool {
	exp1 := regexp.MustCompile(`^\d{1,3}\.\d{1,3}\.\d{1,3}$`)
	return exp1.MatchString(input)
}

func GetEvnVar(input string) (string, bool) {
	reg := regexp.MustCompile(`^\$\{(.*)\}$`)
	if reg.MatchString(input) {
		return reg.FindStringSubmatch(input)[1], true
	}
	return "", false
}

func Check() {

}

func FindStringFirstNum(input string) string {
	reg := regexp.MustCompile(`(\d+(\.\d+)?)`)
	if reg.MatchString(input) {
		return reg.FindStringSubmatch(input)[0]
	}
	return ""
}

func IsDateNumber(input string) bool {
	reg := regexp.MustCompile(`^(\d{6}|\d{8})$`)
	if reg.MatchString(input) {
		return true
	} else {
		return false
	}
}

/***
   b := "rpc error: grpc_code = InvalidArgument desc = 未找到设备 请重试"
     := rpc error: code = Code(1026) desc = 未找到设备
 */
func ParseGrpcError(input string) (code string, desc string, ok bool) {
	reg := regexp.MustCompile(`code = (Code\((\d+)\)|([\w\d]+)) desc = ([^".]*)`)
	//res := reg.FindAllStringSubmatch(b,-1)
	if !reg.MatchString(input) {
		return "", "", false
	}
	res := reg.FindStringSubmatch(input)
	if res[2] != ""{
		return res[2], res[4], true
	}else{
		return res[3], res[4], true
	}
}

//判断ip地址是否是私有ip
func IsLanIp(ip string) bool {
	if strings.HasPrefix(ip, "10.") ||
		strings.HasPrefix(ip, "192.168.") ||
		(strings.Compare(ip, "172.16") >= 0 && strings.Compare(ip, "172.32") < 0) {
		return true
	}
	return false
}