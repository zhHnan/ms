package repo

import (
	"context"
	"hnz.com/ms_serve/ms-user/internal/data/organization"
	"hnz.com/ms_serve/ms-user/internal/database"
)

type OrganizationRepo interface {
	FindOrganizationByMemId(ctx context.Context, memId int64) ([]*organization.Organization, error)
	SaveOrganization(conn database.DBConn, ctx context.Context, org *organization.Organization) error
}
