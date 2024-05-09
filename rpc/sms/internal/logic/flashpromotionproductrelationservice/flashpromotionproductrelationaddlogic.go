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
	err := query.SmsFlashPromotionProductRelation.WithContext(l.ctx).Create(&model.SmsFlashPromotionProductRelation{
		FlashPromotionID:        in.FlashPromotionId,
		FlashPromotionSessionID: in.FlashPromotionSessionId,
		ProductID:               in.ProductId,
		FlashPromotionPrice:     float64(in.FlashPromotionPrice),
		FlashPromotionCount:     in.FlashPromotionCount,
		FlashPromotionLimit:     in.FlashPromotionLimit,
		Sort:                    in.Sort,
	})

	if err != nil {
		return nil, err
	}

	return &smsclient.FlashPromotionProductRelationAddResp{}, nil
}
