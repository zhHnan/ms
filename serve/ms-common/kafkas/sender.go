package kafkas

import (
	"context"
	"errors"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

type KafkaWriter struct {
	w    *kafka.Writer
	data chan LogData
}
type LogData struct {
	Data  []byte
	Topic string
}

func GetWriter(addr string) *KafkaWriter {
	w := &kafka.Writer{
		Addr: kafka.TCP(addr),
		// Balancer 用于选择每条消息的分区.
		Balancer: &kafka.LeastBytes{},
	}
	writer := &KafkaWriter{
		w:    w,
		data: make(chan LogData, 100),
	}
	go writer.sendMsg()
	return writer
}

func (k *KafkaWriter) Send(data LogData) {
	k.data <- data

}
func (k *KafkaWriter) Close() {
	if k.w != nil {
		_ = k.w.Close()
	}
}

func (k *KafkaWriter) sendMsg() {
	for {
		select {
		case data := <-k.data:
			message := []kafka.Message{
				{
					Topic: data.Topic,
					Value: data.Data,
				},
			}
			var err error
			const retries = 3
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			for i := 0; i < retries; i++ {
				err = k.w.WriteMessages(ctx, message...)
				if err == nil {
					break
				}
				if errors.Is(err, kafka.LeaderNotAvailable) || errors.Is(err, context.DeadlineExceeded) {
					time.Sleep(time.Millisecond * 250)
					continue
				}
				if err != nil {
					log.Println("kafka 写入失败：", err.Error())
				}
			}
		}
	}
}
