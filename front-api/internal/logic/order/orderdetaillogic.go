package order

import (
	"context"
	"encoding/json"

	"zero-admin/common/ctxdata"
	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"
	"zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderDetailLogic {
	return &OrderDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderDetailLogic) OrderDetail(req *types.OrderDetailReq) (resp *types.OrderDetailResp, err error) {
	memberId := ctxdata.GetUidFromCtx(l.ctx)
	orderlist, err := l.svcCtx.Oms.OrderByOrderId(l.ctx, &omsclient.OrderByOrderIdReq{
		MemberId: memberId,
		OrderId:  req.OrderId,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("订单列表查询失败,参数：%s,响应：%s", reqStr, err.Error())
		return &types.OrderDetailResp{
			Errno:  1,
			Errmsg: err.Error(),
		}, nil
	}
	order := orderlist.List
	//商品信息
	var goodsListData []types.GoodsListData
	for _, data := range order.GoodsList {
		goodsListData = append(goodsListData, types.GoodsListData{
			Id:             data.Id,
			GoodsName:      data.GoodsName,
			Number:         data.Number,
			PicUrl:         data.PicUrl,
			Specifications: data.Specifications,
			Price:          data.Price,
			Quantity:       data.Quantity,
		})
	}

	//订单信息
	orderDetailData := types.OrderDetailData{
		Id:              order.Id,
		OrderSn:         order.OrderSn,
		ActualPrice:     order.ActualPrice,
		OrderStatusText: order.OrderStatusText,
		HandleOption:    order.HandleOption,
		CreateTime:      order.CreateTime,
		IsDelivery:      order.IsDelivery,
		//商品信息
		GoodsList: goodsListData,
	}

	return &types.OrderDetailResp{
		Errno:  0,
		Data:   orderDetailData,
		Errmsg: "订单列表查询成功",
	}, nil
}
