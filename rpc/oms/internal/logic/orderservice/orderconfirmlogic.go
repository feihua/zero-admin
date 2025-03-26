package orderservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/oms/gen/model"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"
	"time"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// OrderConfirmLogic
/*
Author: LiuFeiHua
Date: 2023/12/7 9:37
*/
type OrderConfirmLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderConfirmLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderConfirmLogic {
	return &OrderConfirmLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// OrderConfirm 确认收货
func (l *OrderConfirmLogic) OrderConfirm(in *omsclient.OrderConfirmReq) (*omsclient.OrderConfirmResp, error) {
	q := query.OmsOrder
	// 1.查询订单是否存在
	order, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.OrderId), q.MemberID.Eq(in.MemberId)).First()

	if err != nil {
		return nil, errors.New("用户订单不存在,确认收货失败")
	}

	// 检测是否能够确认收货
	// 订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单
	// 如果订单已经发货，没有收货，则可收货操作,
	if order.Status != 2 {
		logc.Errorf(l.ctx, "确认收货失败,参数：%+v,订单状态：%d", in, order.Status)
		return nil, errors.New("确认收货失败")
	}

	// 设置订单已确认状态
	item := model.OmsOrder{}
	item.ID = in.OrderId
	item.ConfirmStatus = 1
	item.Status = 3
	now := time.Now()
	item.ReceiveTime = &now
	_, err = q.WithContext(l.ctx).Where(q.ID.Eq(in.OrderId)).Updates(item)
	if err != nil {
		logc.Errorf(l.ctx, "确认收货失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("确认收货失败")
	}

	return &omsclient.OrderConfirmResp{}, nil

}
