package kafkas

import (
	"encoding/json"
	"time"
)

type FieldMap map[string]any
type KafkaLog struct {
	Type string
	// 行为
	Action   string
	Time     string
	Msg      string
	Field    FieldMap
	FuncName string
}

func Error(err error, funcName string, fieldMap FieldMap) []byte {
	kl := KafkaLog{
		Type:     "error",
		Action:   "click",
		Time:     time.Now().Format("2006-01-02 15:04:05"),
		Msg:      err.Error(),
		FuncName: funcName,
		Field:    fieldMap,
	}
	bytes, _ := json.Marshal(kl)
	return bytes
}
func Info(msg string, funcName string, fieldMap FieldMap) []byte {
	kl := KafkaLog{
		Type:     "info",
		Action:   "click",
		Time:     time.Now().Format("2006-01-02 15:04:05"),
		Msg:      msg,
		FuncName: funcName,
		Field:    fieldMap,
	}
	bytes, _ := json.Marshal(kl)
	return bytes
}
