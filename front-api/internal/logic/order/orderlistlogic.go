package order

import (
	"context"
	"encoding/json"
	"zero-admin/common/ctxdata"
	"zero-admin/rpc/oms/omsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) OrderListLogic {
	return OrderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderListLogic) OrderList(req types.OrderListReq) (resp *types.OrderListResp, err error) {

	memberId := ctxdata.GetUidFromCtx(l.ctx)
	orderList, err := l.svcCtx.Oms.OrderListByMemberId(l.ctx, &omsclient.OrderListByMemberIdReq{
		MemberId: memberId,
		Current:  req.Page,
		PageSize: req.Limit,
		Status:   req.Status,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("订单列表查询失败,参数：%s,响应：%s", reqStr, err.Error())
		return &types.OrderListResp{
			Errno:  1,
			Errmsg: err.Error(),
		}, nil
	}

	var orderListData []types.OrderListData

	for _, item := range orderList.List {
		//商品信息
		var goodsListData []types.GoodsListData
		for _, data := range item.GoodsList {
			goodsListData = append(goodsListData, types.GoodsListData{
				Id:             data.Id,
				GoodsName:      data.GoodsName,
				Number:         data.Number,
				PicUrl:         data.PicUrl,
				Specifications: data.Specifications,
				Price:          data.Price,
			})
		}

		//订单信息
		orderListData = append(orderListData, types.OrderListData{
			Id:              item.Id,
			OrderSn:         item.OrderSn,
			ActualPrice:     item.ActualPrice,
			OrderStatusText: item.OrderStatusText,
			HandleOption:    item.HandleOption,
			CreateTime:      item.CreateTime,
			//商品信息
			GoodsList:  goodsListData,
			IsDelivery: item.IsDelivery,
		})
	}

	return &types.OrderListResp{
		Errno:  0,
		Data:   orderListData,
		Errmsg: "订单列表查询成功",
	}, nil
}
