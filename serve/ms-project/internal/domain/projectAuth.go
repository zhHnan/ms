package domain

import (
	"context"
	"strconv"
	"time"

	"go.uber.org/zap"
	"hnz.com/ms_serve/ms-common/errs"
	"hnz.com/ms_serve/ms-project/internal/dao"
	"hnz.com/ms_serve/ms-project/internal/data/account"
	"hnz.com/ms_serve/ms-project/internal/data/node"
	"hnz.com/ms_serve/ms-project/internal/database"
	"hnz.com/ms_serve/ms-project/internal/repo"
	"hnz.com/ms_serve/ms-project/pkg/model"
)

type ProjectAuthDomain struct {
	projectAuthRepo     repo.ProjectAuthRepo
	projectNodeDomain   *ProjectNodeDomain
	projectAuthNodeRepo repo.ProjectAuthNodeRepo
	accountDomain       *AccountDomain
}

func NewProjectAuthDomain() *ProjectAuthDomain {
	return &ProjectAuthDomain{
		projectAuthRepo:     dao.NewProjectAuthDao(),
		projectNodeDomain:   NewProjectNodeDomain(),
		projectAuthNodeRepo: dao.NewProjectAuthNodeDao(),
		accountDomain:       NewAccountDomain(),
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

func (d *ProjectAuthDomain) AuthListPage(orgCode int64, page int64, pageSize int64) ([]*account.ProjectAuthDisplay, int64, *errs.BError) {
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	list, total, err := d.projectAuthRepo.FindAuthListPage(c, orgCode, page, pageSize)
	if err != nil {
		zap.L().Error("project AuthList projectAuthRepo.FindAuthList error", zap.Error(err))
		return nil, 0, model.DataBaseError
	}
	var pdList []*account.ProjectAuthDisplay
	for _, v := range list {
		display := v.ToDisplay()
		pdList = append(pdList, display)
	}
	return pdList, total, nil
}

func (d *ProjectAuthDomain) AllNodeAndAuth(authId int64) ([]*node.ProjectNodeAuthTree, []string, *errs.BError) {
	treeList, err := d.projectNodeDomain.AllNodeList()
	if err != nil {
		return nil, nil, err
	}
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	authNodeList, dbErr := d.projectAuthNodeRepo.FindNodeStringList(c, authId)
	if dbErr != nil {
		return nil, nil, model.DataBaseError
	}
	list := node.ToAuthNodeTreeList(treeList, authNodeList)
	return list, authNodeList, nil
}

func (d *ProjectAuthDomain) Save(conn database.DBConn, authId int64, nodes []string) error {
	err := d.projectAuthNodeRepo.DeleteByAuthId(context.Background(), conn, authId)
	if err != nil {
		return model.DataBaseError
	}

	err = d.projectAuthNodeRepo.Save(context.Background(), conn, authId, nodes)
	if err != nil {
		return model.DataBaseError
	}
	return nil
}

func (d *ProjectAuthDomain) AuthNodes(memberId int64) ([]string, *errs.BError) {
	res, err := d.accountDomain.FindAccount(memberId)
	if err != nil {
		return nil, err
	}
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	authorize := res.Authorize
	authId, _ := strconv.ParseInt(authorize, 10, 64)
	authNodeList, dbErr := d.projectAuthNodeRepo.FindNodeStringList(c, authId)
	if dbErr != nil {
		return nil, model.DataBaseError
	}
	return authNodeList, nil
}
