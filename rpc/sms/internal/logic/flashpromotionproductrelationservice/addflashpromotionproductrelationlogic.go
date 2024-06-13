package flashpromotionproductrelationservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddFlashPromotionProductRelationLogic 添加商品限时购与商品关系表
/*
Author: LiuFeiHua
Date: 2024/6/12 17:40
*/
type AddFlashPromotionProductRelationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddFlashPromotionProductRelationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFlashPromotionProductRelationLogic {
	return &AddFlashPromotionProductRelationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddFlashPromotionProductRelation 添加商品限时购与商品关系表
func (l *AddFlashPromotionProductRelationLogic) AddFlashPromotionProductRelation(in *smsclient.AddFlashPromotionProductRelationReq) (*smsclient.AddFlashPromotionProductRelationResp, error) {
	var data []*model.SmsFlashPromotionProductRelation
	for _, item := range in.Data {
		data = append(data, &model.SmsFlashPromotionProductRelation{
			FlashPromotionID:        item.FlashPromotionId,
			FlashPromotionSessionID: item.FlashPromotionSessionId,
			ProductID:               item.ProductId,
			FlashPromotionPrice:     item.FlashPromotionPrice,
			FlashPromotionCount:     item.FlashPromotionCount,
			FlashPromotionLimit:     item.FlashPromotionLimit,
			Sort:                    item.Sort,
		})
	}

	err := query.SmsFlashPromotionProductRelation.WithContext(l.ctx).Create(data...)

	if err != nil {
		return nil, err
	}

	return &smsclient.AddFlashPromotionProductRelationResp{}, nil
}
