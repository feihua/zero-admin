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

type RoleUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleUpdateLogic {
	return &RoleUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleUpdateLogic) RoleUpdate(in *sys.RoleUpdateReq) (*sys.RoleUpdateResp, error) {
	_ = l.svcCtx.RoleModel.Update(l.ctx, &sysmodel.SysRole{
		Id:         in.Id,
		Name:       in.Name,
		Remark:     sql.NullString{String: in.Remark},
		UpdateBy:   sql.NullString{String: in.LastUpdateBy, Valid: true},
		UpdateTime: sql.NullTime{Time: time.Now()},
		DelFlag:    0,
		Status:     in.Status,
	})

	return &sys.RoleUpdateResp{}, nil
}
