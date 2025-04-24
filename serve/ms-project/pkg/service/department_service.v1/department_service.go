package department_service_v1

import (
	"context"
	"github.com/jinzhu/copier"
	"hnz.com/ms_serve/ms-common/encrypts"
	"hnz.com/ms_serve/ms-common/errs"
	departmentRpc "hnz.com/ms_serve/ms-grpc/department"
	"hnz.com/ms_serve/ms-project/internal/dao"
	"hnz.com/ms_serve/ms-project/internal/database/tran"
	"hnz.com/ms_serve/ms-project/internal/domain"
	"hnz.com/ms_serve/ms-project/internal/repo"
)

type DepartmentService struct {
	departmentRpc.UnimplementedDepartmentServiceServer
	cache            repo.Cache
	transaction      tran.Transaction
	departmentDomain *domain.DepartmentDomain
}

// New 初始化
func New() *DepartmentService {
	return &DepartmentService{
		cache:            dao.Rc,
		transaction:      dao.NewTrans(),
		departmentDomain: domain.NewDepartmentDomain(),
	}
}

func (d *DepartmentService) List(ctx context.Context, msg *departmentRpc.DepartmentReqMessage) (*departmentRpc.ListDepartmentMessage, error) {
	organizationCode := encrypts.DecryptToRes(msg.OrganizationCode)
	var parentDepartmentCode int64
	if msg.ParentDepartmentCode != "" {
		parentDepartmentCode = encrypts.DecryptToRes(msg.ParentDepartmentCode)
	}
	dps, total, err := d.departmentDomain.List(
		organizationCode,
		parentDepartmentCode,
		msg.Page,
		msg.PageSize)
	if err != nil {
		return nil, errs.GrpcError(err)
	}
	var list []*departmentRpc.DepartmentMessage
	_ = copier.Copy(&list, dps)
	return &departmentRpc.ListDepartmentMessage{List: list, Total: total}, nil
}

func (d *DepartmentService) Save(ctx context.Context, msg *departmentRpc.DepartmentReqMessage) (*departmentRpc.DepartmentMessage, error) {
	organizationCode := encrypts.DecryptToRes(msg.OrganizationCode)
	var departmentCode int64
	if msg.DepartmentCode != "" {
		departmentCode = encrypts.DecryptToRes(msg.DepartmentCode)
	}
	var parentDepartmentCode int64
	if msg.ParentDepartmentCode != "" {
		parentDepartmentCode = encrypts.DecryptToRes(msg.ParentDepartmentCode)
	}
	dp, err := d.departmentDomain.Save(organizationCode, departmentCode, parentDepartmentCode, msg.Name)
	if err != nil {
		return &departmentRpc.DepartmentMessage{}, errs.GrpcError(err)
	}
	var res = &departmentRpc.DepartmentMessage{}
	_ = copier.Copy(res, dp)
	return res, nil
}
