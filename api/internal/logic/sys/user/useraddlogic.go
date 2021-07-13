package logic

import (
	"context"
	"encoding/json"
	"go-zero-admin/api/internal/common/errorx"
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
		JobId:    req.JobId,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("添加用户信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("添加用户失败")
	}

	return &types.AddUserResp{
		Code:    "000000",
		Message: "添加用户成功",
	}, nil
}
