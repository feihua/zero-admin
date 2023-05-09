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
		Id:           0,
		Name:         in.Name,
		ParentId:     in.ParentId,
		Url:          sql.NullString{String: in.Url},
		Perms:        sql.NullString{String: in.Perms},
		Type:         in.Type,
		Icon:         sql.NullString{String: in.Icon},
		OrderNum:     sql.NullInt64{Int64: in.OrderNum},
		CreateBy:     in.CreateBy,
		CreateTime:   time.Time{},
		DelFlag:      0,
		VuePath:      sql.NullString{String: in.VuePath},
		VueComponent: sql.NullString{String: in.VueComponent},
		VueIcon:      sql.NullString{String: in.VueIcon},
		VueRedirect:  sql.NullString{String: in.VueRedirect},
	})
	//count, _ := l.svcCtx.UserModel.Count(l.ctx)

	if err != nil {
		return nil, err
	}

	return &sys.MenuAddResp{}, nil
}
