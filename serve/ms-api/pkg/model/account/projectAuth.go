package account

type ProjectAuth struct {
	Id               int64  `json:"id"`
	OrganizationCode string `json:"organization_code"`
	Title            string `json:"title"`
	CreateAt         string `json:"create_at"`
	Sort             int    `json:"sort"`
	Status           int    `json:"status"`
	Desc             string `json:"desc"`
	CreateBy         int64  `json:"create_by"`
	IsDefault        int    `json:"is_default"`
	Type             string `json:"type"`
	CanDelete        int    `json:"canDelete"`
}

type ProjectAuthReq struct {
	Action string `form:"action"`
	Id     int64  `form:"id"`
}

type ProjectNodeAuthTree struct {
	Id       int64                  `json:"id"`
	Node     string                 `json:"node"`
	Title    string                 `json:"title"`
	IsMenu   int                    `json:"is_menu"`
	IsLogin  int                    `json:"is_login"`
	IsAuth   int                    `json:"is_auth"`
	Pnode    string                 `json:"pnode"`
	Key      string                 `json:"key"`
	Checked  bool                   `json:"checked"`
	Children []*ProjectNodeAuthTree `json:"children"`
}
