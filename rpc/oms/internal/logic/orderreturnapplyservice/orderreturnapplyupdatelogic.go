package orderreturnapplyservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"time"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// OrderReturnApplyUpdateLogic 订单退货申请
/*
Author: LiuFeiHua
Date: 2024/5/14 14:53
*/
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

// OrderReturnApplyUpdate 修改订单退货申请状态
func (l *OrderReturnApplyUpdateLogic) OrderReturnApplyUpdate(in *omsclient.OrderReturnApplyUpdateReq) (*omsclient.OrderReturnApplyUpdateResp, error) {

	q := query.OmsOrderReturnApply
	returnApply, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.Id)).First()
	if err != nil {
		return nil, err
	}

	returnApply.Status = in.Status
	//确认退货
	if in.Status == 1 {
		returnApply.CompanyAddressID = in.CompanyAddressId
		returnApply.HandleTime = time.Now()
		returnApply.HandleNote = in.HandleNote
		returnApply.HandleMan = in.HandleMan
		returnApply.ReturnAmount = in.ReturnAmount
	}
	//完成退货
	if in.Status == 2 {
		returnApply.CompanyAddressID = in.CompanyAddressId
		returnApply.ReceiveMan = in.ReceiveMan
		returnApply.ReceiveTime = time.Now()
		returnApply.ReceiveNote = &in.ReceiveNote
	}
	//拒绝退货
	if in.Status == 3 {
		returnApply.HandleTime = time.Now()
		returnApply.HandleNote = in.HandleNote
		returnApply.HandleMan = in.HandleMan
	}

	_, err = q.WithContext(l.ctx).Updates(returnApply)

	return &omsclient.OrderReturnApplyUpdateResp{}, nil
}
