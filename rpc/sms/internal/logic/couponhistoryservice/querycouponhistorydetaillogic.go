package couponhistoryservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryCouponHistoryDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryCouponHistoryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCouponHistoryDetailLogic {
	return &QueryCouponHistoryDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询优惠券使用、领取历史表详情
func (l *QueryCouponHistoryDetailLogic) QueryCouponHistoryDetail(in *smsclient.QueryCouponHistoryDetailReq) (*smsclient.QueryCouponHistoryDetailResp, error) {
	// todo: add your logic here and delete this line

	return &smsclient.QueryCouponHistoryDetailResp{}, nil
}
