package logic

import (
	"context"

	"github.com/feihua/zero-admin/api/web/internal/svc"
	"github.com/feihua/zero-admin/api/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WebLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWebLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WebLogic {
	return &WebLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WebLogic) Web(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
