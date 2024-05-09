package roleservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sys/gen/model"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// RoleAddLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 15:54
*/
type RoleAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleAddLogic {
	return &RoleAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// RoleAdd 新增角色
func (l *RoleAddLogic) RoleAdd(in *sysclient.RoleAddReq) (*sysclient.RoleAddResp, error) {
	err := query.SysRole.WithContext(l.ctx).Create(&model.SysRole{
		Name:     in.Name,
		Remark:   &in.Remark,
		CreateBy: in.CreateBy,
		DelFlag:  1,
		Status:   in.Status,
	})

	if err != nil {
		logc.Errorf(l.ctx, "新增角色信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, err
	}
	return &sysclient.RoleAddResp{}, nil
}
