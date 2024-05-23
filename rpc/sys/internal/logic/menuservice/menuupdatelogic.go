package menuservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sys/gen/model"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// MenuUpdateLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 15:46
*/
type MenuUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuUpdateLogic {
	return &MenuUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// MenuUpdate 更新菜单
func (l *MenuUpdateLogic) MenuUpdate(in *sysclient.MenuUpdateReq) (*sysclient.MenuUpdateResp, error) {
	q := query.SysMenu
	menu, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.Id)).First()
	if err != nil {
		logc.Errorf(l.ctx, "更新菜单信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, err
	}

	_, err = q.WithContext(l.ctx).Updates(&model.SysMenu{
		ID:            in.Id,
		Name:          in.Name,
		ParentID:      in.ParentId,
		URL:           in.Url,
		Perms:         in.Perms,
		Type:          in.Type,
		Icon:          in.Icon,
		OrderNum:      in.OrderNum,
		CreateBy:      menu.CreateBy,
		UpdateBy:      in.UpdateBy,
		DelFlag:       in.DelFlag,
		VuePath:       in.VuePath,
		VueComponent:  in.VueComponent,
		VueIcon:       in.VueIcon,
		VueRedirect:   in.VueRedirect,
		BackgroundURL: in.BackgroundUrl,
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新菜单信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, err
	}

	return &sysclient.MenuUpdateResp{}, nil
}
