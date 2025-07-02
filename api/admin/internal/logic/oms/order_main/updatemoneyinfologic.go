package order_main

import (
	"context"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMoneyInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMoneyInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMoneyInfoLogic {
	return &UpdateMoneyInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateMoneyInfoLogic) UpdateMoneyInfo(req *types.UpdateMoneyInfoReq) (resp *types.BaseResp, err error) {
	// todo: add your logic here and delete this line

	return
}
