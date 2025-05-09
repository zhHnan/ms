package apiProject

type ProjectNodeTree struct {
	Id       int64              `json:"id"`
	Node     string             `json:"node"`
	Title    string             `json:"title"`
	IsMenu   int                `json:"is_menu"`
	IsLogin  int                `json:"is_login"`
	IsAuth   int                `json:"is_auth"`
	Pnode    string             `json:"pnode"`
	Children []*ProjectNodeTree `json:"children"`
}

type ProjectAuthReq struct {
	Action string `form:"action"`
	Id     int64  `form:"id"`
	Nodes  string `form:"nodes" json:"nodes"`
}
