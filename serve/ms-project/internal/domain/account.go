package domain

import (
	"context"
	"fmt"
	"hnz.com/ms_serve/ms-common/encrypts"
	"hnz.com/ms_serve/ms-common/errs"
	"hnz.com/ms_serve/ms-project/internal/dao"
	"hnz.com/ms_serve/ms-project/internal/data/account"
	"hnz.com/ms_serve/ms-project/internal/repo"
	"hnz.com/ms_serve/ms-project/pkg/model"
	"time"
)

type AccountDomain struct {
	accountRepo       repo.AccountRepo
	memberAccountRepo repo.MemberAccountRepo
	userDomain        *UserDomain
	departmentDomain  *DepartmentDomain
}

func NewAccountDomain() *AccountDomain {
	return &AccountDomain{
		accountRepo:       dao.NewAccountDao(),
		memberAccountRepo: dao.NewMemberAccountDao(),
		userDomain:        NewUserDomain(),
		departmentDomain:  NewDepartmentDomain(),
	}
}

func (d *AccountDomain) AccountList(organizationCode string, memberId int64, page int64, pageSize int64, departmentCode string, searchType int32) ([]*account.MemberAccountDisplay, int64, *errs.BError) {
	condition := ""
	organizationCodeId := encrypts.DecryptToRes(organizationCode)
	departmentCodeId := encrypts.DecryptToRes(departmentCode)
	switch searchType {
	case 1:
		condition = "status = 1"
	case 2:
		condition = "department_code = NULL"
	case 3:
		condition = "status = 0"
	case 4:
		condition = fmt.Sprintf("status = 1 and department_code = %d", departmentCodeId)
	default:
		condition = "status = 1"
	}
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	list, total, err := d.memberAccountRepo.FindList(c, condition, organizationCodeId, departmentCodeId, page, pageSize)
	if err != nil {
		return nil, 0, model.DataBaseError
	}
	var dList []*account.MemberAccountDisplay
	for _, v := range list {
		display := v.ToDisplay()
		memberInfo, _ := d.userDomain.MemberInfo(c, v.MemberCode)
		display.Avatar = memberInfo.Avatar
		if v.DepartmentCode > 0 {
			department, err := d.departmentDomain.FindDepartmentById(v.DepartmentCode)
			if err != nil {
				return nil, 0, err
			}
			display.Departments = department.Name
		}
		dList = append(dList, display)
	}
	return dList, total, nil
}

func (d *AccountDomain) FindAccount(memId int64) (*account.MemberAccount, *errs.BError) {
	c, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	memberAccount, err := d.memberAccountRepo.FindByMemberId(c, memId)
	if err != nil {
		return nil, model.DataBaseError
	}
	return memberAccount, nil
}
