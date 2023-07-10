package menuservicelogic

import (
	"context"
	"database/sql"
	"zero-admin/rpc/model/sysmodel"
	"zero-admin/rpc/sys/sysclient"

	"zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *MenuUpdateLogic) MenuUpdate(in *sysclient.MenuUpdateReq) (*sysclient.MenuUpdateResp, error) {
	menu, err := l.svcCtx.MenuModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	err = l.svcCtx.MenuModel.Update(l.ctx, &sysmodel.SysMenu{
		Id:            in.Id,
		Name:          in.Name,
		ParentId:      in.ParentId,
		Url:           sql.NullString{String: in.Url, Valid: true},
		Perms:         sql.NullString{String: in.Perms, Valid: true},
		Type:          in.Type,
		Icon:          sql.NullString{String: in.Icon, Valid: true},
		OrderNum:      sql.NullInt64{Int64: in.OrderNum, Valid: true},
		CreateBy:      menu.CreateBy,
		UpdateBy:      sql.NullString{String: in.LastUpdateBy, Valid: true},
		DelFlag:       in.DelFlag,
		VuePath:       sql.NullString{String: in.VuePath, Valid: true},
		VueComponent:  sql.NullString{String: in.VueComponent, Valid: true},
		VueIcon:       sql.NullString{String: in.VueIcon, Valid: true},
		VueRedirect:   sql.NullString{String: in.VueRedirect, Valid: true},
		BackgroundUrl: in.BackgroundUrl,
	})

	if err != nil {
		return nil, err
	}

	return &sysclient.MenuUpdateResp{}, nil
}
