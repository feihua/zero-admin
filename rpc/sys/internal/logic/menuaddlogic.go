package logic

import (
	"context"
	"database/sql"
	"zero-admin/rpc/model/sysmodel"

	"zero-admin/rpc/sys/internal/svc"
	"zero-admin/rpc/sys/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuAddLogic {
	return &MenuAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MenuAddLogic) MenuAdd(in *sys.MenuAddReq) (*sys.MenuAddResp, error) {
	_, err := l.svcCtx.MenuModel.Insert(l.ctx, &sysmodel.SysMenu{
		Name:         in.Name,
		ParentId:     in.ParentId,
		Url:          sql.NullString{String: in.Url, Valid: true},
		Perms:        sql.NullString{String: in.Perms, Valid: true},
		Type:         in.Type,
		Icon:         sql.NullString{String: in.Icon, Valid: true},
		OrderNum:     sql.NullInt64{Int64: in.OrderNum, Valid: true},
		CreateBy:     in.CreateBy,
		DelFlag:      in.DelFlag,
		VuePath:      sql.NullString{String: in.VuePath, Valid: true},
		VueComponent: sql.NullString{String: in.VueComponent, Valid: true},
		VueIcon:      sql.NullString{String: in.VueIcon, Valid: true},
		VueRedirect:  sql.NullString{String: in.VueRedirect, Valid: true},
	})
	//count, _ := l.svcCtx.UserModel.Count(l.ctx)

	if err != nil {
		return nil, err
	}

	return &sys.MenuAddResp{}, nil
}
