package logic

import (
	"context"
	"go-zero-admin/rpc/sys/sysclient"
	"strconv"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

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
	deptId, _ := strconv.ParseInt(req.DeptId, 10, 64)
	_, err := l.svcCtx.Sys.UserUpdate(l.ctx, &sysclient.UserUpdateReq{
		Id:           req.Id,
		Email:        req.Email,
		Mobile:       req.Mobile,
		Name:         req.Name,
		NickName:     req.NickName,
		Avatar:       req.Avatar,
		DeptId:       deptId,
		LastUpdateBy: "admin",
	})

	if err != nil {
		return nil, err
	}

	return &types.UpdateUserResp{}, nil
}
