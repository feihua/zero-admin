package homerecommendproductservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteHomeRecommendProductLogic 删除人气推荐商品表
/*
Author: LiuFeiHua
Date: 2024/6/12 17:56
*/
type DeleteHomeRecommendProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteHomeRecommendProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteHomeRecommendProductLogic {
	return &DeleteHomeRecommendProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteHomeRecommendProduct 删除人气推荐商品表
func (l *DeleteHomeRecommendProductLogic) DeleteHomeRecommendProduct(in *smsclient.DeleteHomeRecommendProductReq) (*smsclient.DeleteHomeRecommendProductResp, error) {
	q := query.SmsHomeRecommendProduct
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		return nil, err
	}

	return &smsclient.DeleteHomeRecommendProductResp{}, nil
}
