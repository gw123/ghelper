package reg_util

import "regexp"

/***
  ip 格式[10.0.0.1]
*/
func IsIP(input string) bool {
	exp1 := regexp.MustCompile(`^\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}$`)
	return exp1.MatchString(input)
}
