package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/fpay/foundation-go/job"
	"log"
)

const MyJobQueueName = "pos-bind"

type MyJob struct {
	data string `json:"data"`
}

func (m *MyJob) Queue() string {
	return MyJobQueueName
}

func (m *MyJob) Delay() int {
	return 0 // no delay
}

func (m *MyJob) Marshal() ([]byte, error) {
	data := make(map[string]string)
	data["name"] = "hello"
	data["www"] = "www"
	return json.Marshal(data)
}

func main() {
	//manager := job.NewJobManager(job.JobManagerOptions{
	//	"web-server",
	//	"K6VUXh39U87x",
	//	"127.0.0.1",
	//	5672,
	//	"/lh_dev7",
	//})
	manager := job.NewJobManager(job.JobManagerOptions{
		"web-server",
		"K6VUXh39U87x",
		"127.0.0.1",
		25672,
		"/lh_dev7",
	})

	// 发布一个任务
	job := &MyJob{"whatever"}
	err := manager.Dispatch(context.Background(), job)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("发送成功")
}
