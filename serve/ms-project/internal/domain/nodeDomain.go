package domain

import (
	"context"
	"hnz.com/ms_serve/ms-common/errs"
	"hnz.com/ms_serve/ms-project/internal/dao"
	"hnz.com/ms_serve/ms-project/internal/data/node"
	"hnz.com/ms_serve/ms-project/internal/repo"
	"hnz.com/ms_serve/ms-project/pkg/model"
)

type ProjectNodeDomain struct {
	projectNodeRepo repo.ProjectNodeRepo
}

func (d *ProjectNodeDomain) TreeList() ([]*node.ProjectNodeTree, *errs.BError) {
	nodes, err := d.projectNodeRepo.FindAll(context.Background())
	if err != nil {
		return nil, model.DataBaseError
	}
	treeList := node.ToNodeTreeList(nodes)
	return treeList, nil
}

func NewProjectNodeDomain() *ProjectNodeDomain {
	return &ProjectNodeDomain{
		projectNodeRepo: dao.NewProjectNodeDao(),
	}
}
