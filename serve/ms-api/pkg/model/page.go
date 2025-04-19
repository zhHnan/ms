package model

import "github.com/gin-gonic/gin"

type Page struct {
	Page     int64 `json:"page,omitempty" form:"page"`
	PageSize int64 `json:"pageSize,omitempty" form:"pageSize"`
}

func (p *Page) Bind(c *gin.Context) {
	_ = c.ShouldBind(&p)
	if p.PageSize == 0 {
		p.PageSize = 10
	}
	if p.Page == 0 {
		p.Page = 1
	}
}
