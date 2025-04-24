package domain

import (
	"context"
	"hnz.com/ms_serve/ms-project/internal/rpc"

	"hnz.com/ms_serve/ms-grpc/user/login"
)

type UserDomain struct {
	lc login.LoginServiceClient
}

func NewUserDomain() *UserDomain {
	return &UserDomain{
		lc: rpc.UserClient,
	}
}
func (u *UserDomain) MemberList(ctx context.Context, mIdList []int64) ([]*login.MemberMessage, map[int64]*login.MemberMessage, error) {
	messageList, err := rpc.UserClient.FindMemberByIds(ctx, &login.UserMessage{MemberIds: mIdList})
	mMap := make(map[int64]*login.MemberMessage)
	for _, v := range messageList.MemberList {
		mMap[v.Id] = v
	}
	return messageList.MemberList, mMap, err
}
func (d *UserDomain) MemberInfo(c context.Context, mId int64) (*login.MemberMessage, error) {
	message, err := rpc.UserClient.FindMemberInfoById(c, &login.UserMessage{MemId: mId})
	return message, err
}
