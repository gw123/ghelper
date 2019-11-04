package cache

import (
	"github.com/go-redis/redis"
	"reflect"
	"testing"
	"time"
)

type Group struct {
	Title string
}

func TestGetCache(t *testing.T) {
	type args struct {
		client Cache
		key    string
		out    interface{}
		call   func() (interface{}, error)
	}

	option := redis.Options{Addr: "127.0.0.1:6379", DB: 0, Password: "gao123456"}

	g := &Group{}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"loacl",
			args{
				NewRedisCache(redis.NewClient(&option), ""),
				"group:2",
				g,
				func() (i interface{}, e error) {
					return &Group{Title: "title"}, nil
				},
			},
			false,
		},
		{
			"loacl",
			args{
				NewRedisCache(redis.NewClient(&option), ""),
				"group:2",
				nil,
				func() (i interface{}, e error) {
					return &Group{Title: "title"}, nil
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetCache(tt.args.client, tt.args.key, tt.args.out, time.Second*5, tt.args.call); (err != nil) != tt.wantErr {
				t.Errorf("GetCache() error = %v, wantErr %v", err, tt.wantErr)
			}
			if g, ok := tt.args.out.(*Group); ok {
				if g.Title != "title" {
					t.Error(g)
				}
			}
		})
	}
}

func TestNewRedisCache(t *testing.T) {
	option := redis.Options{Addr: "127.0.0.1:6379", DB: 0, Password: "gao123456"}
	redisClient := redis.NewClient(&option)
	cache := NewRedisCache(redisClient, "xyt")

	type args struct {
		client *redis.Client
		prekey string
	}
	tests := []struct {
		name string
		args args
		want *RedisCache
	}{
		{
			"xyt",
			args{
				client: redisClient,
				prekey: "xyt",
			},
			cache,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRedisCache(tt.args.client, tt.args.prekey); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRedisCache() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCache_Del(t *testing.T) {
	option := redis.Options{Addr: "127.0.0.1:6379", DB: 0, Password: "gao123456"}
	redisClient := redis.NewClient(&option)

	type fields struct {
		client *redis.Client
		prekey string
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"xyt",
			fields{
				client: redisClient,
				prekey: "xyt",
			},
			args{key: "key1"},
			false,
		},
		{
			"xyt",
			fields{
				client: redis.NewClient(&redis.Options{Addr: "127.0.0.1:6479", DB: 0, Password: "gao123456"}),
				prekey: "xyt",
			},
			args{key: "key1"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := RedisCache{
				client: tt.fields.client,
				prekey: tt.fields.prekey,
			}
			if err := r.Del(tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("Del() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRedisCache_Get(t *testing.T) {
	type fields struct {
		client *redis.Client
		prekey string
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			"xyt",
			fields{
				client: redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379", DB: 0, Password: "gao123456"}),
				prekey: "xyt",
			},
			args{key: "key1"},
			"hello",
			false,
		},
		{
			"xyt",
			fields{
				client: redis.NewClient(&redis.Options{Addr: "127.0.0.1:6479", DB: 0, Password: "gao123456"}),
				prekey: "xyt",
			},
			args{key: "key1"},
			"hello",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := RedisCache{
				client: tt.fields.client,
				prekey: tt.fields.prekey,
			}
			got, err := r.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisCache_Set(t *testing.T) {
	type fields struct {
		client *redis.Client
		prekey string
	}
	type args struct {
		key        string
		value      interface{}
		expiration time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"xyt",
			fields{
				client: redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379", DB: 0, Password: "gao123456"}),
				prekey: "xyt",
			},
			args{key: "key1", value: "123", expiration: time.Minute},
			false,
		},
		{
			"xyt",
			fields{
				client: redis.NewClient(&redis.Options{Addr: "127.0.0.1:6479", DB: 0, Password: "gao123456"}),
				prekey: "xyt",
			},
			args{key: "key1", value: "123", expiration: time.Minute},
			true,
		},
	}

	cache := NewRedisCache(redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379", DB: 0, Password: "gao123456"}), "xyt")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := RedisCache{
				client: tt.fields.client,
				prekey: tt.fields.prekey,
			}
			if err := r.Set(tt.args.key, tt.args.value, tt.args.expiration); (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			v, err := cache.Get(tt.args.key)
			t.Log(v)
			if err != nil && v.(string) == tt.args.value {
				t.Fail()
			}
		})
	}
}
