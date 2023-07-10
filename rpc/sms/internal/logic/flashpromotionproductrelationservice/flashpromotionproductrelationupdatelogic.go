package flashpromotionproductrelationservicelogic

import (
	"context"
	"zero-admin/rpc/model/smsmodel"
	"zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
	"zero-admin/rpc/sms/internal/svc"
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
