package flashpromotionproductrelationservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FlashPromotionProductRelationAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFlashPromotionProductRelationAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlashPromotionProductRelationAddLogic {
	return &FlashPromotionProductRelationAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FlashPromotionProductRelationAddLogic) FlashPromotionProductRelationAdd(in *smsclient.FlashPromotionProductRelationAddReq) (*smsclient.FlashPromotionProductRelationAddResp, error) {
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

	return &smsclient.FlashPromotionProductRelationAddResp{}, nil
}
