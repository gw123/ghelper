package main

import (
	"fmt"
	"github.com/spf13/viper"
	_  "github.com/spf13/viper/remote"
	"time"
)

//func init() {
//	viper.SetConfigName("config")         // name of config file (without extension)
//	viper.AddConfigPath("/etc/appname/")  // path to look for the config file in
//	viper.AddConfigPath("$HOME/.appname") // call multiple times to add many search paths
//	viper.AddConfigPath(".")              // optionally look for config in the working directory
//	err := viper.ReadInConfig()           // Find and read the config file
//	if err != nil { // Handle errors reading the config file
//		panic(fmt.Errorf("Fatal error config file: %s \n", err))
//	}
//}

func ReadFromEnv() {
	config := viper.New()

	err := config.BindEnv("GW")
	if err != nil {
		panic(err)
	}
	//config.AutomaticEnv()

	path := config.Get("GW")
	fmt.Println(path)
}

type Info struct {
	BindAddr string `json:"bind_addr"`
}

func fromEtcd() {
	var runtime_viper = viper.New()
	runtime_viper.AddRemoteProvider("etcd", "http://192.167.0.172:2379", "/xyt/config.json")
	runtime_viper.SetConfigType("json") // because there is no file extension in a stream of bytes, supported extensions are "json", "toml", "yaml", "yml", "properties", "props", "prop", "env", "dotenv"

	err := runtime_viper.ReadRemoteConfig()

	if err != nil {
		panic(err)
	}
	info := &Info{}
	// unmarshal config
	runtime_viper.Unmarshal(&info)

	// open a goroutine to watch remote changes forever
	go func() {
		for {
			time.Sleep(time.Second * 5) // delay after each request

			// currently, only tested with etcd support
			err := runtime_viper.WatchRemoteConfig()
			if err != nil {
				fmt.Println(err)
				continue
			}
			runtime_viper.Unmarshal(&info)
		}
	}()
}

func main() {
	fromEtcd()
}
