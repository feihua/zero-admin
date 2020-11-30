package logic

import (
	"context"
	"go-zero-admin/rpc/sys/sysclient"
	"strconv"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

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
	deptId, _ := strconv.ParseInt(req.DeptId, 10, 64)
	_, err := l.svcCtx.Sys.UserAdd(l.ctx, &sysclient.UserAddReq{
		Email:    req.Email,
		Mobile:   req.Mobile,
		Name:     req.Name,
		NickName: req.NickName,
		DeptId:   deptId,
		CreateBy: "admin",
	})

	if err != nil {
		return nil, err
	}

	return &types.AddUserResp{}, nil
}
