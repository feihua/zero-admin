package flashpromotionproductrelationservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteFlashPromotionProductRelationLogic 删除商品限时购与商品关系表
/*
Author: LiuFeiHua
Date: 2024/6/12 17:40
*/
type DeleteFlashPromotionProductRelationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteFlashPromotionProductRelationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFlashPromotionProductRelationLogic {
	return &DeleteFlashPromotionProductRelationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteFlashPromotionProductRelation 删除商品限时购与商品关系表
func (l *DeleteFlashPromotionProductRelationLogic) DeleteFlashPromotionProductRelation(in *smsclient.DeleteFlashPromotionProductRelationReq) (*smsclient.DeleteFlashPromotionProductRelationResp, error) {
	q := query.SmsFlashPromotionProductRelation
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...), q.FlashPromotionID.Eq(in.FlashPromotionId)).Delete()

	if err != nil {
		return nil, err
	}

	return &smsclient.DeleteFlashPromotionProductRelationResp{}, nil
}
