package project_service_v1

import (
	task_service_v1 "hnz.com/ms_serve/ms-grpc/task"
	"hnz.com/ms_serve/ms-project/internal/dao"
	"hnz.com/ms_serve/ms-project/internal/database/tran"
	"hnz.com/ms_serve/ms-project/internal/repo"
)

type TaskService struct {
	task_service_v1.UnimplementedTaskServiceServer
	cache       repo.Cache
	transaction tran.Transaction
}

// 初始化
func New() *TaskService {
	return &TaskService{
		cache:       dao.Rc,
		transaction: dao.NewTrans(),
	}
}
