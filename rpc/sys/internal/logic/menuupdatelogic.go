package logic

import (
	"context"
	"database/sql"
	"time"
	"zero-admin/rpc/model/sysmodel"

	"zero-admin/rpc/sys/internal/svc"
	"zero-admin/rpc/sys/sys"

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

func (l *MenuUpdateLogic) MenuUpdate(in *sys.MenuUpdateReq) (*sys.MenuUpdateResp, error) {
	menu, err := l.svcCtx.MenuModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	err = l.svcCtx.MenuModel.Update(l.ctx, &sysmodel.SysMenu{
		Id:            in.Id,
		Name:          in.Name,
		ParentId:      in.ParentId,
		Url:           sql.NullString{String: in.Url},
		Perms:         sql.NullString{String: in.Perms},
		Type:          in.Type,
		Icon:          sql.NullString{String: in.Icon},
		OrderNum:      sql.NullInt64{Int64: in.OrderNum},
		CreateBy:      menu.CreateBy,
		CreateTime:    menu.CreateTime,
		UpdateBy:      sql.NullString{String: in.LastUpdateBy},
		UpdateTime:    sql.NullTime{Time: time.Now()},
		DelFlag:       0,
		VuePath:       sql.NullString{String: in.VuePath},
		VueComponent:  sql.NullString{String: in.VueComponent},
		VueIcon:       sql.NullString{String: in.VueIcon},
		VueRedirect:   sql.NullString{String: in.VueRedirect},
		BackgroundUrl: "",
	})

	if err != nil {
		return nil, err
	}

	return &sys.MenuUpdateResp{}, nil
}
