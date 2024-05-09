package flashpromotionproductrelationservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type FlashPromotionProductRelationUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFlashPromotionProductRelationUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlashPromotionProductRelationUpdateLogic {
	return &FlashPromotionProductRelationUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FlashPromotionProductRelationUpdateLogic) FlashPromotionProductRelationUpdate(in *smsclient.FlashPromotionProductRelationUpdateReq) (*smsclient.FlashPromotionProductRelationUpdateResp, error) {
	q := query.SmsFlashPromotionProductRelation
	_, err := q.WithContext(l.ctx).Updates(&model.SmsFlashPromotionProductRelation{
		ID:                      in.Id,
		FlashPromotionID:        in.FlashPromotionId,
		FlashPromotionSessionID: in.FlashPromotionSessionId,
		ProductID:               in.ProductId,
		FlashPromotionPrice:     in.FlashPromotionPrice,
		FlashPromotionCount:     in.FlashPromotionCount,
		FlashPromotionLimit:     in.FlashPromotionLimit,
		Sort:                    in.Sort,
	})

	if err != nil {
		return nil, err
	}

	return &smsclient.FlashPromotionProductRelationUpdateResp{}, nil
}
