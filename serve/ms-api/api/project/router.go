package project

import (
	"github.com/gin-gonic/gin"
	"hnz.com/ms_serve/ms-api/api/midd"
	"hnz.com/ms_serve/ms-api/api/rpc"
	"hnz.com/ms_serve/ms-api/router"
	"log"
)

type RouterUser struct {
}

func init() {
	log.Println("init user-api routers...")
	ru := &RouterUser{}
	router.Register(ru)
}

func (*RouterUser) Route(r *gin.Engine) {
	rpc.InitProjectRpc()
	h := New()
	group := r.Group("/project")
	group.Use(midd.TokenVerify())
	group.POST("/index", h.index)
	group.POST("/project/selfList", h.projectList)
	group.POST("/project", h.projectList)
	group.POST("/project_template", h.projectTemplate)
	group.POST("/project/save", h.projectSave)
	group.POST("/project/read", h.projectRead)
	group.POST("/project/recycle", h.projectRecycle)
	group.POST("/project/recovery", h.projectRecovery)
	group.POST("/project_collect/collect", h.projectCollect)
	group.POST("/project/edit", h.projectEdit)
	group.POST("/project/getLogBySelfProject", h.getLogBySelfProject)

	t := NewTask()
	group.POST("/task_stages", t.taskStages)
	group.POST("/project_member/index", t.taskMemberList)
	group.POST("/task_stages/tasks", t.taskList)
	group.POST("/task/save", t.taskSave)
	group.POST("/task/sort", t.taskSort)
	group.POST("/task/selfList", t.selfList)
	group.POST("/task/read", t.taskRead)
	group.POST("/task_member", t.listTaskMember)
	group.POST("/task/taskLog", t.taskLog)
	group.POST("/task/_taskWorkTimeList", t.taskWorkTimeList)
	group.POST("/task/saveTaskWorkTime", t.saveTaskWorkTime)
	group.POST("/file/uploadFiles", t.uploadFiles)
	group.POST("/task/taskSources", t.taskSources)
	group.POST("/task/createComment", t.createComment)

	a := NewAccount()
	group.POST("/account", a.account)

	d := NewDepartment()
	group.POST("/department", d.department)
	group.POST("/department/save", d.save)
	group.POST("/department/read", d.read)
}
