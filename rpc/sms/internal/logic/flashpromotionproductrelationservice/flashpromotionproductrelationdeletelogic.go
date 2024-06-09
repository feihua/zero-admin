package flashpromotionproductrelationservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// FlashPromotionProductRelationDeleteLogic 限时购与产品关糸
/*
Author: LiuFeiHua
Date: 2024/5/8 10:13
*/
type FlashPromotionProductRelationDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFlashPromotionProductRelationDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlashPromotionProductRelationDeleteLogic {
	return &FlashPromotionProductRelationDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// FlashPromotionProductRelationDelete 删除限时购与产品关糸
func (l *FlashPromotionProductRelationDeleteLogic) FlashPromotionProductRelationDelete(in *smsclient.FlashPromotionProductRelationDeleteReq) (*smsclient.FlashPromotionProductRelationDeleteResp, error) {
	q := query.SmsFlashPromotionProductRelation
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...), q.FlashPromotionID.Eq(in.FlashPromotionId)).Delete()

	if err != nil {
		return nil, err
	}

	return &smsclient.FlashPromotionProductRelationDeleteResp{}, nil
}
