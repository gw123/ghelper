package cache

import (
	"encoding/json"
	"errors"
	"time"
)

//统计没有命中缓存个数 ,调用GetCache最好通过
const CacheStatisticsPrefix = "Statistics:"

var openStatistics bool = true

func CloseStatistics() {
	openStatistics = false
}

func OpenStatistics() {
	openStatistics = true
}

type Cache interface {
	Get(key string) (interface{}, error)
	Set(key string, value interface{}, expiration time.Duration) error
	Del(key string) error
	Incr(key string) error
}

//先从缓存读取数据,如果不存在调用 call方法获取后在存放到数据库中
func GetCache(client Cache, key string, out interface{}, expiration time.Duration, call func() (interface{}, error)) ( error) {
	if out == nil {
		return errors.New("out can not  be nil")
	}
	isOK := false
	val, err := client.Get(key)
	//是否需要重新刷新缓存
	if err == nil && val != "" {
		if v, ok := val.(string); ok {
			err := json.Unmarshal([]byte(v), out)
			if err != nil {
				isOK = false
			} else {
				isOK = true
			}
		}
	}

	if val == "" || !isOK {
		{
			//记录没有命中的缓存次数, 方便后面优化代码
			client.Incr(CacheStatisticsPrefix + key)
		}
		newVal, err := call()
		if err != nil {
			return err
		}
		data, err := json.Marshal(newVal)
		if err != nil {
			return err
		}
		if err := client.Set(key, data, expiration); err != nil {
			return err
		}
		err = json.Unmarshal(data, out)
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}
