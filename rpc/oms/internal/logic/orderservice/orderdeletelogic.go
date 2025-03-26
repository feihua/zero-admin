package orderservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/oms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderDeleteLogic {
	return &OrderDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderDeleteLogic) OrderDelete(in *omsclient.OrderDeleteReq) (*omsclient.OrderDeleteResp, error) {

	q := query.OmsOrder
	// 1.查询订单是否存在
	order, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.OrderId), q.MemberID.Eq(in.MemberId)).First()

	if err != nil {
		return nil, errors.New("用户订单不存在,删除用户订单失败")
	}

	// 检测是否能够删除订单
	// 订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单
	// 如果订单已经取消或是已完成，则可删除
	if order.Status == 3 || order.Status == 4 || order.Status == 5 {
		_, err = q.WithContext(l.ctx).Where(q.ID.Eq(in.OrderId)).Delete()
		if err != nil {
			logc.Errorf(l.ctx, "删除用户订单失败,参数:%+v,异常:%s", in, err.Error())
			return nil, errors.New("删除用户订单失败")
		}
		return &omsclient.OrderDeleteResp{}, nil
	}

	logc.Errorf(l.ctx, "删除用户订单失败,参数：%s,订单状态：%s", in, order.Status)
	return nil, errors.New("删除用户订单失败")

}
