package map_tool

import (
	"reflect"
	"fmt"
	"sort"
)

func Dump(item interface{}) {
	itemType := reflect.TypeOf(item)
	itemValue := reflect.ValueOf(item)
	fieldNum := itemType.NumField()
	for i := 0; i < fieldNum; i++ {
		field := itemType.Field(i)
		fmt.Printf("%s: %v ,%v \n", field.Name, field.Type, field.Anonymous)
		firstByte := []byte(field.Name)[0]
		if firstByte >= 'A' && firstByte <= 'Z' {
			//私有字段不能访问问题
			//panic: reflect.Value.Interface: cannot return value obtained from unexported field or method
			fieldVal := itemValue.Field(i).Interface()
			fmt.Printf("[ %v ]\n", fieldVal)
		}
	}
}

func GetStringFieldMap(item interface{}) map[string]string {
	itemType := reflect.TypeOf(item)
	itemValue := reflect.ValueOf(item)
	fieldNum := itemType.NumField()
	retVal := make(map[string]string, 0)
	for i := 0; i < fieldNum; i++ {
		field := itemType.Field(i)
		firstByte := []byte(field.Name)[0]
		if firstByte >= 'A' && firstByte <= 'Z' && field.Type.Name() == "string" {
			fieldVal := itemValue.Field(i).Interface()
			retVal[field.Name] = fieldVal.(string)
		}
	}
	return retVal
}

/***
 同名字段处理:后面的覆盖前面
 */
func MergeField(item1 interface{}, items ...interface{}) map[string]string {
	itemsLen := len(items)
	if itemsLen == 0 {
		return GetStringFieldMap(item1)
	}
	retMap := GetStringFieldMap(item1)
	for index := 0; index < itemsLen; index++ {
		tempMap := GetStringFieldMap(items[index])

		for key, val := range tempMap {
			retMap[key] = val
		}
	}
	return retMap
}

/***
 同名字段处理: 前面的覆盖后面面
 */
func MergeField2(item1 interface{}, items ...interface{}) map[string]string {
	itemsLen := len(items)
	if itemsLen == 0 {
		return GetStringFieldMap(item1)
	}
	retMap := GetStringFieldMap(item1)
	for index := 0; index < itemsLen; index++ {
		tempMap := GetStringFieldMap(items[index])

		for key, val := range tempMap {
			if _, ok := retMap[key]; !ok {
				retMap[key] = val
			}
		}
	}
	return retMap
}

/***
  排序map的 keys
 */
func SortMapKeys(item map[string]string) []string {
	keys := make([]string, 0)
	for key, _ := range item {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

/***
	排序map的 values
 */
func SortMapValues(item map[string]string) []string {
	values := make([]string, 0)
	for _, value := range item {
		values = append(values, value)
	}
	sort.Strings(values)
	return values
}
