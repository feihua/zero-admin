package role

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateRoleStatusLogic 更新角色状态
/*
Author: LiuFeiHua
Date: 2024/5/29 17:59
*/
type UpdateRoleStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateRoleStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleStatusLogic {
	return &UpdateRoleStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateRoleStatus 更新角色状态
func (l *UpdateRoleStatusLogic) UpdateRoleStatus(req *types.UpdateRoleStatusReq) (resp *types.UpdateRoleStatusResp, err error) {
	roleReq := sysclient.UpdateRoleStatusReq{
		Ids:        req.RoleIds,    // 角色编号
		RoleStatus: req.RoleStatus, // 角色状态
		UpdateBy:   l.ctx.Value("userName").(string),
	}
	_, err = l.svcCtx.RoleService.UpdateRoleStatus(l.ctx, &roleReq)

	if err != nil {
		logc.Errorf(l.ctx, "更新角色状态失败,参数:%+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.UpdateRoleStatusResp{
		Code:    "000000",
		Message: "更新角色状态成功",
	}, nil
}
