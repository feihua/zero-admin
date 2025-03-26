package homerecommendproductservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateHomeRecommendProductStatusLogic 更新人气推荐商品状态
/*
Author: LiuFeiHua
Date: 2024/6/12 17:57
*/
type UpdateHomeRecommendProductStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateHomeRecommendProductStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateHomeRecommendProductStatusLogic {
	return &UpdateHomeRecommendProductStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateHomeRecommendProductStatus 更新人气推荐商品状态
func (l *UpdateHomeRecommendProductStatusLogic) UpdateHomeRecommendProductStatus(in *smsclient.UpdateHomeRecommendProductStatusReq) (*smsclient.UpdateHomeRecommendProductStatusResp, error) {
	q := query.SmsHomeRecommendProduct
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.RecommendStatus, in.RecommendStatus)

	if err != nil {
		logc.Errorf(l.ctx, "更新人气推荐商品状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新人气推荐商品状态失败")
	}

	return &smsclient.UpdateHomeRecommendProductStatusResp{}, nil
}
