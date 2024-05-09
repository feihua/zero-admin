package flashpromotionservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

// FlashPromotionDeleteLogic 限时购列表
/*
Author: LiuFeiHua
Date: 2024/5/8 10:14
*/
type FlashPromotionDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFlashPromotionDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlashPromotionDeleteLogic {
	return &FlashPromotionDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// FlashPromotionDelete 删除限时购列表
func (l *FlashPromotionDeleteLogic) FlashPromotionDelete(in *smsclient.FlashPromotionDeleteReq) (*smsclient.FlashPromotionDeleteResp, error) {
	q := query.SmsFlashPromotion
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		return nil, err
	}

	return &smsclient.FlashPromotionDeleteResp{}, nil
}
