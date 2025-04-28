package role

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

// DeleteRoleLogic 删除角色
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

// DeleteRole 删除角色
func (l *DeleteRoleLogic) DeleteRole(req *types.DeleteRoleReq) (*types.BaseResp, error) {

	_, err := l.svcCtx.RoleService.DeleteRole(l.ctx, &sysclient.DeleteRoleReq{
		Ids: req.Ids, // 角色id
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据roleIds: %+v,删除角色异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return res.Success()
}
