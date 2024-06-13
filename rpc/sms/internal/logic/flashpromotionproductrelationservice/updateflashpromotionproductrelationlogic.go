package flashpromotionproductrelationservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateFlashPromotionProductRelationLogic 更新商品限时购与商品关系表
/*
Author: LiuFeiHua
Date: 2024/6/12 17:41
*/
type UpdateFlashPromotionProductRelationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateFlashPromotionProductRelationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFlashPromotionProductRelationLogic {
	return &UpdateFlashPromotionProductRelationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateFlashPromotionProductRelation 更新商品限时购与商品关系表
func (l *UpdateFlashPromotionProductRelationLogic) UpdateFlashPromotionProductRelation(in *smsclient.UpdateFlashPromotionProductRelationReq) (*smsclient.UpdateFlashPromotionProductRelationResp, error) {
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

	return &smsclient.UpdateFlashPromotionProductRelationResp{}, nil
}
