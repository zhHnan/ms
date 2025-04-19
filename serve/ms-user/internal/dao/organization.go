package dao

import (
	"context"
	"hnz.com/ms_serve/ms-user/internal/data/organization"
	"hnz.com/ms_serve/ms-user/internal/database"
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
func (o *OrganizationDao) FindOrganizationByMemId(ctx context.Context, memId int64) ([]*organization.Organization, error) {
	var orgs []*organization.Organization
	err := o.conn.Session(ctx).Where("member_id = ?", memId).Find(&orgs).Error
	return orgs, err
}

// SaveOrganization 保存组织信息
func (o *OrganizationDao) SaveOrganization(conn database.DBConn, ctx context.Context, org *organization.Organization) error {
	o.conn = conn.(*gorms.GormConn)
	return o.conn.Tx(ctx).Create(org).Error
}
