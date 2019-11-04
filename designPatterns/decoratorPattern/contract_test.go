package decoratorPattern

import (
	"github.com/go-redis/redis"
	"github.com/gw123/ghelper/cache"
	"testing"
)

//测试装饰器模式 , 实现在获取数据同时将结果更新到缓存中
func TestNewResourceCacheService(t *testing.T) {
	dbService := NewResourceDbService("xytschool")
	cache := cache.NewRedisCache(redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379", DB: 0, Password: "gao123456"}), "test_")
	cacheService := NewResourceCacheService(dbService, cache)
	out, err := cacheService.GetGroup(102)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(out)

	out, err = cacheService.GetGroup(102)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(out)
}

