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

// DeleteRoleLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 15:37
*/
type DeleteRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteRoleLogic {
	return DeleteRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteRole 删除角色(id为1的是系统预留超级管理员角色,不能删除)
func (l *DeleteRoleLogic) DeleteRole(req *types.DeleteRoleReq) (*types.DeleteRoleResp, error) {

	_, err := l.svcCtx.RoleService.RoleDelete(l.ctx, &sysclient.RoleDeleteReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据roleId: %+v,删除角色异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.DeleteRoleResp{
		Code:    "000000",
		Message: "删除角色成功",
	}, nil
}
