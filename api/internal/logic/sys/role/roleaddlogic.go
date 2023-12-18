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

// RoleAddLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 15:34
*/
type RoleAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) RoleAddLogic {
	return RoleAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// RoleAdd 新增角色
func (l *RoleAddLogic) RoleAdd(req types.AddRoleReq) (*types.AddRoleResp, error) {
	roleAddReq := sysclient.RoleAddReq{
		Name:     req.Name,
		Remark:   req.Remark,
		CreateBy: l.ctx.Value("userName").(string),
		Status:   req.Status,
	}

	_, err := l.svcCtx.RoleService.RoleAdd(l.ctx, &roleAddReq)

	if err != nil {
		logc.Errorf(l.ctx, "添加角色信息失败,参数:%+v,异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("添加角色失败")
	}

	return &types.AddRoleResp{
		Code:    "000000",
		Message: "添加角色成功",
	}, nil
}
