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
		Avatar:     req.Avatar,
		DeptId:     req.DeptId,
		Email:      req.Email,
		Id:         req.Id,
		Mobile:     req.Mobile,
		NickName:   req.NickName,
		Remark:     req.Remark,
		UpdateBy:   l.ctx.Value("userName").(string),
		UserName:   req.UserName,
		UserStatus: req.UserStatus,
		PostIds:    req.PostIds,
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
