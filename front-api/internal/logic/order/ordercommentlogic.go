package order

import (
	"context"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) OrderCommentLogic {
	return OrderCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderCommentLogic) OrderComment(req types.OrderCommentReq) (resp *types.OrderCommentResp, err error) {
	// todo: add your logic here and delete this line

	return
}
