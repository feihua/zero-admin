package user

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateUserLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 14:05
*/
type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateUserLogic {
	return UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateUser 更新用户信息
func (l *UpdateUserLogic) UpdateUser(req *types.UpdateUserReq) (*types.UpdateUserResp, error) {
	userUpdateReq := sysclient.UserUpdateReq{
		Id:         req.Id,
		Email:      req.Email,
		Mobile:     req.Mobile,
		UserName:   req.Name,
		NickName:   req.NickName,
		DeptId:     req.DeptId,
		UpdateBy:   l.ctx.Value("userName").(string),
		RoleId:     req.RoleId,
		UserStatus: req.Status,
		JobId:      req.JobId,
	}

	if _, err := l.svcCtx.UserService.UserUpdate(l.ctx, &userUpdateReq); err != nil {
		logc.Errorf(l.ctx, "更新用户信息失败,参数:%+v,异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("更新用户失败")
	}

	return &types.UpdateUserResp{
		Code:    "000000",
		Message: "更新用户成功",
	}, nil
}
