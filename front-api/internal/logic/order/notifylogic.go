package order

import (
	"context"
	"net/http"

	"github.com/feihua/zero-admin/front-api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

// NotifyLogic
/*
Author: LiuFeiHua
Date: 2023/12/14 16:11
*/
type NotifyLogic struct {
	logx.Logger
	ctx     context.Context
	svcCtx  *svc.ServiceContext
	writer  http.ResponseWriter
	request *http.Request
}

func NewNotifyLogic(w http.ResponseWriter, r *http.Request, ctx context.Context, svcCtx *svc.ServiceContext) *NotifyLogic {
	return &NotifyLogic{Logger: logx.WithContext(ctx), ctx: ctx, svcCtx: svcCtx, writer: w, request: r}
}

// Notify 支付回调用处理
func (l *NotifyLogic) Notify() {
	NewPaymentOperationsUtils(l.ctx, l.svcCtx).AliPayNotify(l.writer, l.request)
}
