package domain

import (
	"context"
	"fmt"
	"hnz.com/ms_serve/ms-common/errs"
	"hnz.com/ms_serve/ms-common/kafkas"
	"hnz.com/ms_serve/ms-project/config"
	"hnz.com/ms_serve/ms-project/internal/dao"
	"hnz.com/ms_serve/ms-project/internal/repo"
	"hnz.com/ms_serve/ms-project/pkg/model"
)

type TaskDomain struct {
	taskRepo repo.TaskRepo
}

func NewTaskDomain() *TaskDomain {
	return &TaskDomain{
		taskRepo: dao.NewTaskDao(),
	}
}

func (d *TaskDomain) FindProjectIdByTaskId(taskId int64) (int64, bool, *errs.BError) {
	fmt.Println("FindProjectIdByTaskId")
	config.SendLog(kafkas.Info("Find", "FindProjectIdByTaskId", kafkas.FieldMap{
		"taskId": taskId,
	}))
	task, err := d.taskRepo.FindTaskById(context.Background(), taskId)
	if err != nil {
		config.SendLog(kafkas.Error(err, "FindProjectIdByTaskId", kafkas.FieldMap{
			"taskId": taskId,
		}))
		return 0, false, model.DataBaseError
	}
	if task == nil {
		return 0, false, nil
	}
	return task.ProjectCode, true, nil
}
