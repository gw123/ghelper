package decoratorPattern

import (
	"fmt"
	"github.com/gw123/ghelper/cache"
	"time"
)

type Group struct {
	Id    int
	Title string
	Desc  string
}

type ResourceService interface {
	GetGroup(id int) (*Group, error)
}

//db
type ResourceDbService struct {
	db string
}

//db获取数据
func NewResourceDbService(db string) *ResourceDbService {
	return &ResourceDbService{db: db}
}

func (r *ResourceDbService) GetGroup(id int) (*Group, error) {
	return &Group{id, "title", "desc"}, nil
}

//缓存
type ResourceCacheService struct {
	inner ResourceService
	cache cache.Cache
}

func NewResourceCacheService(inner ResourceService, cache2 cache.Cache) *ResourceCacheService {
	return &ResourceCacheService{inner: inner, cache: cache2}
}

func (r *ResourceCacheService) GetGroup(id int) (*Group, error) {
	group := &Group{}
	key := fmt.Sprintf("group:%d", id)
	err := cache.GetCache(r.cache, key, group, time.Hour, func() (i interface{}, e error) {
		return r.inner.GetGroup(id)
	})
	if err != nil {
		return nil, err
	}
	return group, nil
}

