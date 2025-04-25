package domain

import (
	"context"
	"hnz.com/ms_serve/ms-common/errs"
	"hnz.com/ms_serve/ms-project/internal/dao"
	"hnz.com/ms_serve/ms-project/internal/data/menu"
	"hnz.com/ms_serve/ms-project/internal/repo"
	"hnz.com/ms_serve/ms-project/pkg/model"
	"time"
)

type MenuDomain struct {
	menuRepo repo.MenuRepo
}

func NewMenuDomain() *MenuDomain {
	return &MenuDomain{
		menuRepo: dao.NewMenuDao(),
	}
}
func (d *MenuDomain) MenuTreeList() ([]*menu.ProjectMenuChild, *errs.BError) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	menus, err := d.menuRepo.FindAll(ctx)
	if err != nil {
		return nil, model.DataBaseError
	}
	menuChildren := menu.CovertChild(menus)
	return menuChildren, nil
}
