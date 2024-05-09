package orderreturnapplyservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"time"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderReturnApplyUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderReturnApplyUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderReturnApplyUpdateLogic {
	return &OrderReturnApplyUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderReturnApplyUpdateLogic) OrderReturnApplyUpdate(in *omsclient.OrderReturnApplyUpdateReq) (*omsclient.OrderReturnApplyUpdateResp, error) {

	q := query.OmsOrderReturnApply
	returnApply, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.Id)).First()
	if err != nil {
		return nil, err
	}

	// 退货中处理退货中的就要设置公司地址
	if in.Status == 1 {
		returnApply.CompanyAddressID = in.CompanyAddressId
	}

	returnApply.Status = in.Status
	returnApply.HandleTime = time.Now()
	returnApply.HandleNote = in.HandleNote
	returnApply.HandleMan = in.HandleMan
	returnApply.ReceiveMan = in.ReceiveMan
	returnApply.ReceiveTime = time.Now()
	returnApply.ReceiveNote = &in.ReceiveNote

	_, err = q.WithContext(l.ctx).Updates(returnApply)

	return &omsclient.OrderReturnApplyUpdateResp{}, nil
}
