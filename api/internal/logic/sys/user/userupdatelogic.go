package logic

import (
	"context"
	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"
	"go-zero-admin/rpc/sys/sysclient"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserUpdateLogic {
	return UserUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserUpdateLogic) UserUpdate(req types.UpdateUserReq) (*types.UpdateUserResp, error) {
	_, err := l.svcCtx.Sys.UserUpdate(l.ctx, &sysclient.UserUpdateReq{
		Id:           req.Id,
		Email:        req.Email,
		Mobile:       req.Mobile,
		Name:         req.Name,
		NickName:     req.NickName,
		DeptId:       req.DeptId,
		LastUpdateBy: "admin",
		RoleId:       req.RoleId,
		Status:       req.Status,
		JobId:        req.JobId,
	})

	if err != nil {
		return nil, err
	}

	return &types.UpdateUserResp{
		Code:    "000000",
		Message: "",
	}, nil
}
