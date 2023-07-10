package logic

import (
	"context"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/sys/sysclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) RoleDeleteLogic {
	return RoleDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleDeleteLogic) RoleDelete(req types.DeleteRoleReq) (*types.DeleteRoleResp, error) {
	_, err := l.svcCtx.RoleService.RoleDelete(l.ctx, &sysclient.RoleDeleteReq{
		Ids: req.Ids,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据roleId: %d,删除角色异常:%s", req.Ids, err.Error())
		return nil, errorx.NewDefaultError("删除角色失败")
	}

	return &types.DeleteRoleResp{
		Code:    "000000",
		Message: "删除角色成功",
	}, nil
}
