package homerecommendproductservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// HomeRecommendProductDeleteLogic 人气商品推荐
/*
Author: LiuFeiHua
Date: 2024/5/8 10:19
*/
type HomeRecommendProductDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomeRecommendProductDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeRecommendProductDeleteLogic {
	return &HomeRecommendProductDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// HomeRecommendProductDelete 删除人气商品推荐
func (l *HomeRecommendProductDeleteLogic) HomeRecommendProductDelete(in *smsclient.HomeRecommendProductDeleteReq) (*smsclient.HomeRecommendProductDeleteResp, error) {
	q := query.SmsHomeRecommendProduct
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		return nil, err
	}

	return &smsclient.HomeRecommendProductDeleteResp{}, nil
}
