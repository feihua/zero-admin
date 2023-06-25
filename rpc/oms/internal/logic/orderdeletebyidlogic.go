package logic

import (
	"context"
	"encoding/json"
	"errors"

	"zero-admin/rpc/oms/internal/svc"
	"zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderDeleteByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderDeleteByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderDeleteByIdLogic {
	return &OrderDeleteByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderDeleteByIdLogic) OrderDeleteById(in *omsclient.OrderDeleteByIdReq) (*omsclient.OrderDeleteResp, error) {
	order, err := l.svcCtx.OmsOrderModel.FindOne(l.ctx, in.Ids[0])

	if err != nil {
		return nil, err
	}

	//检测是否能够删除订单
	//订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单
	//如果订单已经取消或是已完成，则可删除
	if order.Status == 3 || order.Status == 4 || order.Status == 5 {
		if l.svcCtx.OmsOrderModel.Delete(l.ctx, in.Ids[0]) != nil {
			return nil, errors.New("删除用户订单失败")
		}
		return &omsclient.OrderDeleteResp{}, nil
	}

	reqStr, _ := json.Marshal(in)
	logx.WithContext(l.ctx).Errorf("删除用户订单失败,参数：%s,订单状态：%s", reqStr, order.Status)
	return nil, errors.New("删除用户订单失败")
}
