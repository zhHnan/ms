package kafkas

import (
	"encoding/json"
	"testing"
	"time"
)

func TestProducer(t *testing.T) {
	w := GetWriter("127.0.0.1:9092")
	m := make(map[string]string)
	m["projectCode"] = "1001"
	bytes, _ := json.Marshal(m)
	w.Send(LogData{
		Data:  bytes,
		Topic: "msproject_log",
	})
	time.Sleep(2 * time.Second)
}

func TestConsumer(t *testing.T) {
	GetReader([]string{"127.0.0.1:9092"}, "group1", "test")
	for {

	}
}
