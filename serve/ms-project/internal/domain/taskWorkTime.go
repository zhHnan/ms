package domain

import (
	"context"
	"hnz.com/ms_serve/ms-project/internal/data/common"
	"hnz.com/ms_serve/ms-project/internal/data/task"
	"time"

	"go.uber.org/zap"
	"hnz.com/ms_serve/ms-common/errs"
	"hnz.com/ms_serve/ms-project/internal/dao"
	"hnz.com/ms_serve/ms-project/internal/repo"
	"hnz.com/ms_serve/ms-project/pkg/model"
)

type TaskWorkTimeDomain struct {
	taskWorkTimeRepo repo.TaskWorkTimeRepo
	UserDomain       *UserDomain
}

func NewTaskWorkTimeDomain() *TaskWorkTimeDomain {
	return &TaskWorkTimeDomain{
		taskWorkTimeRepo: dao.NewTaskWorkTimeDao(),
		UserDomain:       NewUserDomain(),
	}
}

func (t *TaskWorkTimeDomain) TaskWorkTimeList(taskCode int64) ([]*task.TaskWorkTimeDisplay, *errs.BError) {
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	var list []*task.TaskWorkTime
	var err error
	list, err = t.taskWorkTimeRepo.FindWorkTimeList(c, taskCode)
	if err != nil {
		zap.L().Error("project task TaskWorkTimeList taskWorkTimeRepo.FindWorkTimeList error", zap.Error(err))
		return nil, model.DataBaseError
	}
	if len(list) == 0 {
		return []*task.TaskWorkTimeDisplay{}, nil
	}
	var displayList []*task.TaskWorkTimeDisplay
	var mIdList []int64
	for _, v := range list {
		mIdList = append(mIdList, v.MemberCode)
	}
	_, mMap, err := t.UserDomain.MemberList(c, mIdList)
	if err != nil {
		return nil, errs.ToBError(err)
	}
	for _, v := range list {
		display := v.ToDisplay()
		message := mMap[v.MemberCode]
		m := common.Member{}
		m.Name = message.Name
		m.Id = message.Id
		m.Avatar = message.Avatar
		m.Code = message.Code
		display.Member = m
		displayList = append(displayList, display)
	}
	return displayList, nil
}
