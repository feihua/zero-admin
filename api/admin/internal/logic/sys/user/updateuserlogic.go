package user

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateUserLogic 更新用户信息
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
	userUpdateReq := sysclient.UpdateUserReq{
		Id:         req.Id,         // 编号
		UserName:   req.UserName,   // 用户名
		NickName:   req.NickName,   // 昵称
		Avatar:     req.Avatar,     // 头像
		Password:   req.Password,   // 密码
		Email:      req.Email,      // 邮箱
		Mobile:     req.Mobile,     // 手机号
		UserStatus: req.UserStatus, // 帐号状态（0正常 1停用）
		DeptId:     req.DeptId,     // 部门id
		Remark:     req.Remark,     // 备注信息
		PostIds:    req.PostIds,    // 部门id
		UpdateBy:   l.ctx.Value("userName").(string),
	}

	if _, err := l.svcCtx.UserService.UpdateUser(l.ctx, &userUpdateReq); err != nil {
		logc.Errorf(l.ctx, "更新用户信息失败,参数:%+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.UpdateUserResp{
		Code:    "000000",
		Message: "更新用户成功",
	}, nil
}
