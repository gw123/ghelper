package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/fpay/foundation-go/job"
	"log"
	"time"
)

const MyJobQueueName = "device_pos_bind"
const DeviceLogType_AdminBind = "admin_bind"
const DeviceLogType_BossBind = "boss_bind"
const TimeFormat = "2006-01-02 15:04:05"

type Time struct {
	time.Time
}

func NewTime(t time.Time) *Time {
	return &Time{t}
}

func (t *Time) MarshalJSON() ([]byte, error) {
	return []byte("\"" + t.Time.Format(TimeFormat) + "\""), nil
}

type PosBindJob struct {
	CreatedAt *Time   `json:"created_at"`
	DeviceNo  string `json:"device_no"`
	Meid      int    `json:"meid"`
	Maid      int    `json:"maid"`
	Wuid      int    `json:"wuid"`
	Type      string `json:"type"`
	FactoryId int    `json:"factory_id"`
}

func NewPosBindJob(deviceNo, bindType string, meid, maid, wuid, factory_id int) *PosBindJob {
	return &PosBindJob{
		CreatedAt: NewTime(time.Now()),
		DeviceNo:  deviceNo,
		Meid:      meid,
		Maid:      maid,
		FactoryId: factory_id,
		Type:      bindType,
		Wuid:      wuid,
	}
}

func (m *PosBindJob) Queue() string {
	return "device_pos_bind"
}

func (m *PosBindJob) Delay() int {
	return 0 // no delay
}

func (m *PosBindJob) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

type JobManager interface {
	Dispatch(ctx context.Context, bindJob *PosBindJob) error
}

func main() {
	manager := job.NewJobManager(job.JobManagerOptions{
		"web-server",
		"K6VUXh39U87x",
		"127.0.0.1",
		25672,
		"/lh_dev9",
	})
	// 发布一个任务
	job := NewPosBindJob("LD195EBQ731807", DeviceLogType_AdminBind, 460, 123, 0, 3)


	err := manager.Dispatch(context.Background(), job)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("发送成功")

	// 发布一个任务
	job = NewPosBindJob("LD195EBQ731807", DeviceLogType_BossBind, 460, 0, 123, 3)
	err = manager.Dispatch(context.Background(), job)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("发送成功")
}
