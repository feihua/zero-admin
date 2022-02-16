package logic

import (
	"context"
	"encoding/json"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"
	"zero-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
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
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("更新用户信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("更新用户失败")
	}

	return &types.UpdateUserResp{
		Code:    "000000",
		Message: "更新用户成功",
	}, nil
}
