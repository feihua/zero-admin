package logic

import (
	"context"
	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"
	"go-zero-admin/rpc/sys/sysclient"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserAddLogic {
	return UserAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserAddLogic) UserAdd(req types.AddUserReq) (*types.AddUserResp, error) {
	_, err := l.svcCtx.Sys.UserAdd(l.ctx, &sysclient.UserAddReq{
		Email:    req.Email,
		Mobile:   req.Mobile,
		Name:     req.Name,
		NickName: req.NickName,
		DeptId:   req.DeptId,
		CreateBy: "admin",
		RoleId:   req.RoleId,
	})

	if err != nil {
		return nil, err
	}

	return &types.AddUserResp{
		Code:    "000000",
		Message: "",
	}, nil
}
