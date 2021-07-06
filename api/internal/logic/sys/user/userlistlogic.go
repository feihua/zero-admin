package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"go-zero-admin/api/internal/common/errorx"
	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"
	"go-zero-admin/rpc/sys/sysclient"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserListLogic {
	return UserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserListLogic) UserList(req types.ListUserReq) (*types.ListUserResp, error) {

	userListReq := sysclient.UserListReq{}
	_ = copier.Copy(&userListReq, &req)
	resp, err := l.svcCtx.Sys.UserList(l.ctx, &userListReq)

	if err != nil {
		return nil, errorx.NewDefaultError("查询失败")
	}

	var list []*types.ListUserData

	for _, item := range resp.List {
		listUserData := types.ListUserData{}
		_ = copier.Copy(&listUserData, &item)
		list = append(list, &listUserData)
	}

	return &types.ListUserResp{
		Code:     "000000",
		Message:  "查询用户列表成功",
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
	}, nil
}
