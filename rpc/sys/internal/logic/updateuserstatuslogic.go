package logic

import (
	"context"
	"time"
	"zero-admin/rpc/model/sysmodel"

	"zero-admin/rpc/sys/internal/svc"
	"zero-admin/rpc/sys/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserStatusLogic {
	return &UpdateUserStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserStatusLogic) UpdateUserStatus(in *sys.UserStatusReq) (*sys.UserStatusResp, error) {
	_ = l.svcCtx.UserModel.Update(sysmodel.SysUser{
		Id:             in.Id,
		Status:         in.Status,
		LastUpdateBy:   in.LastUpdateBy,
		LastUpdateTime: time.Now(),
	})

	return &sys.UserStatusResp{}, nil
}
