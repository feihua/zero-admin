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
		ID:                      in.Id,                      // 编号
		FlashPromotionID:        in.FlashPromotionId,        // 限时购id
		FlashPromotionSessionID: in.FlashPromotionSessionId, // 编号
		ProductID:               in.ProductId,               // 商品id
		FlashPromotionPrice:     in.FlashPromotionPrice,     // 限时购价格
		FlashPromotionCount:     in.FlashPromotionCount,     // 限时购数量
		FlashPromotionLimit:     in.FlashPromotionLimit,     // 每人限购数量
		Sort:                    in.Sort,                    // 排序
	})

	if err != nil {
		return nil, err
	}

	return &smsclient.UpdateFlashPromotionProductRelationResp{}, nil
}
