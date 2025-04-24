package domain

import (
	"context"
	"go.uber.org/zap"
	"hnz.com/ms_serve/ms-common/errs"
	"hnz.com/ms_serve/ms-project/internal/dao"
	"hnz.com/ms_serve/ms-project/internal/data/account"
	"hnz.com/ms_serve/ms-project/internal/repo"
	"hnz.com/ms_serve/ms-project/pkg/model"
	"time"
)

type ProjectAuthDomain struct {
	accountRepo       repo.AccountRepo
	memberAccountRepo repo.MemberAccountRepo
	projectAuthRepo   repo.ProjectAuthRepo
}

func NewProjectAuthDomain() *ProjectAuthDomain {
	return &ProjectAuthDomain{
		accountRepo:       dao.NewAccountDao(),
		memberAccountRepo: dao.NewMemberAccountDao(),
		projectAuthRepo:   dao.NewProjectAuthDao(),
	}
}

func (d *ProjectAuthDomain) AuthList(orgCode int64) ([]*account.ProjectAuthDisplay, *errs.BError) {
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	list, err := d.projectAuthRepo.FindAuthList(c, orgCode)
	if err != nil {
		zap.L().Error("project AuthList projectAuthRepo.FindAuthList error", zap.Error(err))
		return nil, model.DataBaseError
	}
	var pdList []*account.ProjectAuthDisplay
	for _, v := range list {
		display := v.ToDisplay()
		pdList = append(pdList, display)
	}
	return pdList, nil
}
