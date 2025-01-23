package menuservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateMenuStatusLogic 更新菜单状态
/*
Author: LiuFeiHua
Date: 2024/5/30 11:36
*/
type UpdateMenuStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMenuStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMenuStatusLogic {
	return &UpdateMenuStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateMenuStatus 更新菜单状态
// 1.判断菜单是否存在
// 2.更新菜单状态
func (l *UpdateMenuStatusLogic) UpdateMenuStatus(in *sysclient.UpdateMenuStatusReq) (*sysclient.UpdateMenuStatusResp, error) {
	sysMenu := query.SysMenu
	q := sysMenu.WithContext(l.ctx)

	// 1.判断菜单是否存在
	count, err := q.Where(sysMenu.ID.Eq(in.Id)).Count()

	if err != nil {
		return nil, errors.New("查询菜单失败")
	}

	if count == 0 {
		return nil, errors.New("菜单不存在")
	}

	// 2.更新菜单状态
	_, err = q.Where(sysMenu.ID.Eq(in.Id)).Update(sysMenu.MenuStatus, in.MenuStatus)

	if err != nil {
		logc.Errorf(l.ctx, "更新菜单状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新菜单状态失败")
	}

	return &sysclient.UpdateMenuStatusResp{}, nil
}
