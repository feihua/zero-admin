package orderpromotionservicelogic

import (
	"context"
	"errors"

	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// QueryOrderPromotionListLogic 查询订单优惠信息列表
/*
Author: LiuFeiHua
Date: 2025/07/03 09:32:56
*/
type QueryOrderPromotionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryOrderPromotionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryOrderPromotionListLogic {
	return &QueryOrderPromotionListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryOrderPromotionList 查询订单优惠信息列表
func (l *QueryOrderPromotionListLogic) QueryOrderPromotionList(in *omsclient.QueryOrderPromotionListReq) (*omsclient.QueryOrderPromotionListResp, error) {
	orderPromotion := query.OmsOrderPromotion
	q := orderPromotion.WithContext(l.ctx)
	if in.OrderId != 2 {
		q = q.Where(orderPromotion.OrderID.Eq(in.OrderId))
	}
	if len(in.OrderNo) > 0 {
		q = q.Where(orderPromotion.OrderNo.Like("%" + in.OrderNo + "%"))
	}
	if in.PromotionType != 2 {
		q = q.Where(orderPromotion.PromotionType.Eq(in.PromotionType))
	}
	if in.PromotionId != 2 {
		q = q.Where(orderPromotion.PromotionID.Eq(in.PromotionId))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询订单优惠信息列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询订单优惠信息列表失败")
	}

	var list []*omsclient.OrderPromotionListData

	for _, item := range result {
		data := omsclient.OrderPromotionListData{
			Id:             item.ID,                              // 主键ID
			OrderId:        item.OrderID,                         // 订单ID
			OrderNo:        item.OrderNo,                         // 订单编号
			PromotionType:  item.PromotionType,                   // 优惠类型：1-优惠券，2-积分抵扣，3-会员折扣，4-促销活动
			PromotionName:  item.PromotionName,                   // 优惠名称
			DiscountAmount: float32(item.DiscountAmount),         // 优惠金额
			CreateTime:     time_util.TimeToStr(item.CreateTime), //

		}
		if item.PromotionID != nil {
			data.PromotionId = *item.PromotionID
		}
		list = append(list, &data)
	}

	return &omsclient.QueryOrderPromotionListResp{
		Total: count,
		List:  list,
	}, nil
}
