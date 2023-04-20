package order

import (
	"context"
	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderGoodsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) OrderGoodsLogic {
	return OrderGoodsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderGoodsLogic) OrderGoods(req types.OrderGoodsReq) (resp *types.OrderGoodsResp, err error) {
	// memberId := ctxdata.GetUidFromCtx(l.ctx)
	// orderList, err := l.svcCtx.Oms.order(l.ctx, &omsclient.OrderListByMemberIdReq{
	// 	MemberId: memberId,
	// 	Current:  req.Page,
	// 	PageSize: req.Limit,
	// })

	// if err != nil {
	// 	reqStr, _ := json.Marshal(req)
	// 	logx.WithContext(l.ctx).Errorf("订单列表查询失败,参数：%s,响应：%s", reqStr, err.Error())
	// 	return &types.OrderListResp{
	// 		Errno:  1,
	// 		Errmsg: err.Error(),
	// 	}, nil
	// }

	// var orderListData []types.OrderListData

	// for _, item := range orderList.List {
	// 	//商品信息
	// 	var goodsListData []types.GoodsListData
	// 	for _, data := range item.GoodsList {
	// 		goodsListData = append(goodsListData, types.GoodsListData{
	// 			Id:             data.Id,
	// 			GoodsName:      data.GoodsName,
	// 			Number:         data.Number,
	// 			PicUrl:         data.PicUrl,
	// 			Specifications: data.Specifications,
	// 			Price:          data.Price,
	// 		})
	// 	}

	// 	//订单信息
	// 	orderGoodsData = append(orderListData, types.OrderGoodsData{
	// 		Id:              item.Id,
	// 		OrderSn:         item.OrderSn,
	// 		ActualPrice:     item.ActualPrice,
	// 		OrderStatusText: item.OrderStatusText,
	// 		HandleOption:    item.HandleOption,
	// 		//商品信息
	// 		GoodsList: goodsListData,
	// 	})
	// }

	// return &types.OrderGoodsResp{
	// 	Errno:  0,
	// 	Data:   orderListData,
	// 	Errmsg: "",
	// }, nil

	return
}
