package user

import (
	"context"
	"github.com/feihua/zero-admin/api/internal/common/errorx"
	"github.com/feihua/zero-admin/api/internal/svc"
	"github.com/feihua/zero-admin/api/internal/types"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// UserListLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 14:04
*/
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

// UserList 查询用户列表信息
func (l *UserListLogic) UserList(req types.ListUserReq) (*types.ListUserResp, error) {

	userListReq := sysclient.UserListReq{}
	_ = copier.Copy(&userListReq, &req)
	resp, err := l.svcCtx.UserService.UserList(l.ctx, &userListReq)

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询用户列表异常:%s", req, err.Error())
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
