package menuservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteMenuLogic 删除菜单
/*
Author: LiuFeiHua
Date: 2023/12/18 15:45
*/
type DeleteMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMenuLogic {
	return &DeleteMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteMenu 删除菜单
// 1.查询菜单是否有子菜单
// 2.查询菜单是否菜单已分配
func (l *DeleteMenuLogic) DeleteMenu(in *sysclient.DeleteMenuReq) (*sysclient.DeleteMenuResp, error) {
	q := query.SysMenu

	// 1.查询菜单是否有子菜单
	count, err := q.WithContext(l.ctx).Where(q.ParentID.Eq(in.Id)).Count()
	if err != nil {
		logc.Errorf(l.ctx, "查询菜单是否有子菜单,异常:%s", err.Error())
		return nil, errors.New(fmt.Sprintf("查删除菜单失败"))
	}

	if count > 0 {
		return nil, errors.New(fmt.Sprintf("存在子菜单,不允许删除"))
	}

	// 2.查询菜单是否菜单已分配
	count, err = query.SysRoleMenu.WithContext(l.ctx).Where(query.SysRoleMenu.MenuID.Eq(in.Id)).Count()
	if err != nil {
		logc.Errorf(l.ctx, "查询菜单是否菜单已分配,异常:%s", err.Error())
		return nil, errors.New(fmt.Sprintf("查删除菜单失败"))
	}

	if count > 0 {
		return nil, errors.New(fmt.Sprintf("菜单已分配,不允许删除"))
	}

	_, err = q.WithContext(l.ctx).Where(q.ID.Eq(in.Id)).Delete()
	if err != nil {
		logc.Errorf(l.ctx, "删除菜单失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除菜单失败")
	}

	return &sysclient.DeleteMenuResp{}, nil
}
