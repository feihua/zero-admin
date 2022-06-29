package logic

import (
	"context"
	"encoding/json"
	"errors"
	"zero-admin/rpc/oms/omsclient"

	"zero-admin/rpc/oms/internal/svc"
	"zero-admin/rpc/oms/oms"

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

func (l *OrderDeleteLogic) OrderDelete(in *oms.OrderDeleteReq) (*oms.OrderDeleteResp, error) {

	order, err := l.svcCtx.OmsOrderModel.FindOne(in.OrderId)

	if err != nil {
		return nil, err
	}

	if order.MemberId != in.UserId {
		return nil, errors.New("用户订单不存在,删除用户订单失败")
	}

	//检测是否能够删除订单
	//订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单
	//如果订单已经取消或是已完成，则可删除
	if order.Status == 3 || order.Status == 4 || order.Status == 5 {
		if l.svcCtx.OmsOrderModel.Delete(in.OrderId) != nil {
			return nil, errors.New("删除用户订单失败")
		}
		return &omsclient.OrderDeleteResp{}, nil
	}

	reqStr, _ := json.Marshal(in)
	logx.WithContext(l.ctx).Errorf("删除用户订单失败,参数：%s,订单状态：%s", reqStr, order.Status)
	return nil, errors.New("删除用户订单失败")

}
