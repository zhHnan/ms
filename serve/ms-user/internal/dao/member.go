package dao

import (
	"context"
	"hnz.com/ms_serve/ms-user/internal/data/member"
	"hnz.com/ms_serve/ms-user/internal/database/gorms"
)

type MemberDao struct {
	conn *gorms.GormConn
}

func NewMemberDao() *MemberDao {
	return &MemberDao{
		conn: gorms.New(),
	}
}

func (m *MemberDao) GetMemberByEmail(ctx context.Context, email string) (bool, error) {
	var count int64
	err := m.conn.Session(ctx).Model(&member.Member{}).Where("email = ?", email).Count(&count).Error
	return count > 0, err
}

func (m *MemberDao) GetMemberByAccount(ctx context.Context, name string) (bool, error) {
	var count int64
	err := m.conn.Session(ctx).Model(&member.Member{}).Where("account = ?", name).Count(&count).Error
	return count > 0, err
}

func (m *MemberDao) GetMemberByMobile(ctx context.Context, mobile string) (bool, error) {
	var count int64
	err := m.conn.Session(ctx).Model(&member.Member{}).Where("mobile = ?", mobile).Count(&count).Error
	return count > 0, err
}
func (m *MemberDao) SaveMember(ctx context.Context, mem *member.Member) error {
	return m.conn.Session(ctx).Create(mem).Error
}
