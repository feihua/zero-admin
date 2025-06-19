package pay

import (
	"context"
	"net/http"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

// NotifyLogic 支付回调用处理
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
func (l *NotifyLogic) Notify() error {
	NewPaymentOperationsUtils(l.ctx, l.svcCtx).AliPayNotify(l.writer, l.request)

	return nil
}
