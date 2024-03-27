package flashpromotionproductrelationservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/model/smsmodel"
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
	err := l.svcCtx.SmsFlashPromotionProductRelationModel.Update(l.ctx, &smsmodel.SmsFlashPromotionProductRelation{
		Id:                      in.Id,
		FlashPromotionId:        in.FlashPromotionId,
		FlashPromotionSessionId: in.FlashPromotionSessionId,
		ProductId:               in.ProductId,
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
