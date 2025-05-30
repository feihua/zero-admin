package user

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/common/res"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateUserRoleListLogic 分配用户角色
/*
Author: LiuFeiHua
Date: 2024/5/23 17:33
*/
type UpdateUserRoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserRoleListLogic {
	return &UpdateUserRoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateUserRoleList 分配用户角色
func (l *UpdateUserRoleListLogic) UpdateUserRoleList(req *types.UpdateUserRoleReq) (resp *types.BaseResp, err error) {
	_, err = l.svcCtx.UserService.UpdateUserRoleList(l.ctx, &sysclient.UpdateUserRoleListReq{
		UserId:  req.UserId,  // 用户id
		RoleIds: req.RoleIds, // 角色id
	})

	if err != nil {
		logc.Errorf(l.ctx, "分配用户角色失败,参数:%+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return res.Success()
}
