package role

import (
	"context"
	"github.com/zeromicro/go-zero/core/logc"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/sys/sysclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// RoleUpdateLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 15:41
*/
type RoleUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) RoleUpdateLogic {
	return RoleUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// RoleUpdate 更新角色(id为1的是系统预留超级管理员角色,不能更新)
func (l *RoleUpdateLogic) RoleUpdate(req types.UpdateRoleReq) (*types.UpdateRoleResp, error) {
	roleUpdateReq := sysclient.RoleUpdateReq{
		Id:           req.Id,
		Name:         req.Name,
		Remark:       req.Remark,
		LastUpdateBy: l.ctx.Value("userName").(string),
		Status:       req.Status,
	}
	_, err := l.svcCtx.RoleService.RoleUpdate(l.ctx, &roleUpdateReq)

	if err != nil {
		logc.Errorf(l.ctx, "更新角色信息失败,参数:%+v,异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("更新角色失败")
	}

	return &types.UpdateRoleResp{
		Code:    "000000",
		Message: "更新角色成功",
	}, nil
}
