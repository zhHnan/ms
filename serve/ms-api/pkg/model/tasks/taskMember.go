package tasks

type TaskMember struct {
	Id                int    `json:"id"`
	Name              string `json:"name"`
	Avatar            string `json:"avatar"`
	Code              string `json:"code"`
	MemberAccountCode string `json:"member_account_code"`
	IsExecutor        int    `json:"is_executor"`
	IsOwner           int    `json:"is_owner"`
}
