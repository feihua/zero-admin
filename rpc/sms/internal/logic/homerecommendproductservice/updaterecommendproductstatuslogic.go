package homerecommendproductservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateRecommendProductStatusLogic 人气商品推荐
/*
Author: LiuFeiHua
Date: 2024/5/14 9:34
*/
type UpdateRecommendProductStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRecommendProductStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRecommendProductStatusLogic {
	return &UpdateRecommendProductStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateRecommendProductStatus 批量修改推荐状态
func (l *UpdateRecommendProductStatusLogic) UpdateRecommendProductStatus(in *smsclient.UpdateRecommendProductStatusReq) (*smsclient.UpdateRecommendProductStatusResp, error) {
	q := query.SmsHomeRecommendProduct
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.RecommendStatus, in.RecommendStatus)

	if err != nil {
		return nil, err
	}

	return &smsclient.UpdateRecommendProductStatusResp{}, nil
}
