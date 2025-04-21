package repo

import (
	"context"
	"hnz.com/ms_serve/ms-project/internal/data/project"
	"hnz.com/ms_serve/ms-project/internal/database"
)

type ProjectRepo interface {
	FindProjectByMemId(ctx context.Context, memId int64, condition string, page int64, size int64) ([]*project.ProjectAndMember, int64, error)
	FindCollectProjectByMemId(ctx context.Context, id int64, page int64, size int64) ([]*project.ProjectAndMember, int64, error)
	SaveProject(conn database.DBConn, ctx context.Context, pr *project.Project) error
	SaveProjectMember(conn database.DBConn, ctx context.Context, pm *project.ProjectMember) error
	FindProjectByPIdAndMemId(ctx context.Context, code int64, id int64) (*project.ProjectAndMember, error)
	FindCollectByPidAndMemId(ctx context.Context, code int64, id int64) (bool, error)
	UpdateDeletedProject(ctx context.Context, code int64, deleted bool) error
	SaveProjectCollect(ctx context.Context, pc *project.ProjectCollection) error
	DeleteProjectCollect(ctx context.Context, memId int64, projectCode int64) error
	UpdateProject(ctx context.Context, proj *project.Project) error
	FindMemberByProjectId(ctx context.Context, projectCode int64) (list []*project.ProjectMember, total int64, err error)
}
