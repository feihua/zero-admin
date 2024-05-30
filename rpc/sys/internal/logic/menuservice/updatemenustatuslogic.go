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
func (l *UpdateMenuStatusLogic) UpdateMenuStatus(in *sysclient.UpdateMenuStatusReq) (*sysclient.UpdateMenuStatusResp, error) {
	q := query.SysMenu.WithContext(l.ctx)
	// 1.根据菜单名称查询菜单是否已存在
	_, err := q.Where(query.SysMenu.ID.In(in.Ids...)).Update(query.SysMenu.MenuStatus, in.MenuStatus)

	if err != nil {
		logc.Errorf(l.ctx, "更新菜单状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新菜单状态失败")
	}

	return &sysclient.UpdateMenuStatusResp{}, nil
}
