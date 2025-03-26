package flashpromotionproductrelationservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// QueryFlashPromotionProductRelationDetailLogic 查询商品限时购与商品关系详情
/*
Author: LiuFeiHua
Date: 2025/01/23 16:18:56
*/
type QueryFlashPromotionProductRelationDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryFlashPromotionProductRelationDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryFlashPromotionProductRelationDetailLogic {
	return &QueryFlashPromotionProductRelationDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryFlashPromotionProductRelationDetail 查询商品限时购与商品关系详情
func (l *QueryFlashPromotionProductRelationDetailLogic) QueryFlashPromotionProductRelationDetail(in *smsclient.QueryFlashPromotionProductRelationDetailReq) (*smsclient.QueryFlashPromotionProductRelationDetailResp, error) {
	item, err := query.SmsFlashPromotionProductRelation.WithContext(l.ctx).Where(query.SmsFlashPromotionProductRelation.ID.Eq(in.ProductId)).First()

	if err != nil {
		logc.Errorf(l.ctx, "查询商品限时购与商品关系详情失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询商品限时购与商品关系详情失败")
	}

	data := &smsclient.QueryFlashPromotionProductRelationDetailResp{
		Id:                      item.ID,                      // 编号
		FlashPromotionId:        item.FlashPromotionID,        // 限时购id
		FlashPromotionSessionId: item.FlashPromotionSessionID, // 编号
		ProductId:               item.ProductID,               // 商品id
		FlashPromotionPrice:     item.FlashPromotionPrice,     // 限时购价格
		FlashPromotionCount:     item.FlashPromotionCount,     // 限时购数量
		FlashPromotionLimit:     item.FlashPromotionLimit,     // 每人限购数量
		Sort:                    item.Sort,                    // 排序
	}

	return data, nil
}
