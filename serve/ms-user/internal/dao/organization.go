package dao

import (
	"context"
	"hnz.com/ms_serve/ms-user/internal/data/organization"
	"hnz.com/ms_serve/ms-user/internal/database/gorms"
)

type OrganizationDao struct {
	conn *gorms.GormConn
}

func NewOrganizationDao() *OrganizationDao {
	return &OrganizationDao{
		conn: gorms.New(),
	}
}

// implements
// FindOrganizationByMemId 根据会员ID查询组织信息
func (o *OrganizationDao) FindOrganizationByMemId(ctx context.Context, memId int64) ([]organization.Organization, error) {
	//TODO implement me
	panic("implement me")
}

// SaveOrganization 保存组织信息
func (o *OrganizationDao) SaveOrganization(ctx context.Context, org *organization.Organization) error {
	return o.conn.Session(ctx).Create(org).Error
}
