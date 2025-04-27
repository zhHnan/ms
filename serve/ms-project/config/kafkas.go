package config

import (
	"context"
	"go.uber.org/zap"
	"hnz.com/ms_serve/ms-common/kafkas"
	"hnz.com/ms_serve/ms-project/internal/dao"
	"hnz.com/ms_serve/ms-project/internal/repo"
	"time"
)

var kafkaWriter *kafkas.KafkaWriter

func InitKafkaWriter() func() {
	kafkaWriter = kafkas.GetWriter("localhost:9092")
	return kafkaWriter.Close
}

func SendLog(data []byte) {
	kafkaWriter.Send(kafkas.LogData{
		Data:  data,
		Topic: "msproject_log",
	})
}

func SendCache(data []byte) {
	kafkaWriter.Send(kafkas.LogData{
		Data:  data,
		Topic: "msproject_cache",
	})
}

type KafkaCache struct {
	R     *kafkas.KafkaReader
	cache repo.Cache
}

func NewKafkaCache() *KafkaCache {
	return &KafkaCache{
		R:     kafkas.GetReader([]string{"localhost:9092"}, "cache_group", "msproject_cache"),
		cache: dao.Rc,
	}
}

func (c *KafkaCache) DeleteCache() {
	for {
		message, err := c.R.R.ReadMessage(context.Background())
		if err != nil {
			zap.L().Error("kafka DeleteCache err", zap.Error(err))
			continue
		}
		if "task" == string(message.Value) {
			fields, err := c.cache.HKeys(context.Background(), string(message.Key))
			if err != nil {
				zap.L().Error("kafka DeleteCache err", zap.Error(err))
				continue
			}
			time.Sleep(1 * time.Second)
			c.cache.Delete(context.Background(), fields)
		}
	}
}
