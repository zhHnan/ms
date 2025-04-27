package kafkas

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type KafkaReader struct {
	R *kafka.Reader
}

func GetReader(brokers []string, groupId, topic string) *KafkaReader {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		// group: 同一个组下的消费者协同工作， 共同消费topic队列的内容
		// 如果不设置GroupID，则每个消费者都会独立消费topic队列的内容
		GroupID:  groupId,
		Topic:    topic,
		MinBytes: 10e3,
		MaxBytes: 10e6,
	})
	reader := &KafkaReader{R: r}
	//go reader.readMsg()
	return reader
}

func (r *KafkaReader) readMsg() {
	for {
		m, err := r.R.ReadMessage(context.Background())
		if err != nil {
			zap.L().Error("kafka receiver read msg err", zap.Error(err))
			continue
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}
}
func (r *KafkaReader) Close() {
	r.R.Close()
}
