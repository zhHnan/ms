package config

import "hnz.com/ms_serve/ms-common/kafkas"

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
