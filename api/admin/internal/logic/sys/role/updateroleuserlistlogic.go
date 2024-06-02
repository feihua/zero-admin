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

// UpdateRoleUserListLogic 更新角色用户
/*
Author: LiuFeiHua
Date: 2024/6/02 17:59
*/
type UpdateRoleUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateRoleUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleUserListLogic {
	return &UpdateRoleUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateRoleUserList 更新角色用户
func (l *UpdateRoleUserListLogic) UpdateRoleUserList(req *types.UpdateRoleUserListReq) (resp *types.UpdateRoleUserListResp, err error) {
	_, err = l.svcCtx.RoleService.UpdateRoleUserList(l.ctx, &sysclient.UpdateRoleUserListReq{
		RoleId:  req.RoleId,
		UserIds: req.UserIds,
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新角色用户失败,参数:%+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.UpdateRoleUserListResp{
		Code:    "000000",
		Message: "更新角色用户成功",
	}, nil
}
