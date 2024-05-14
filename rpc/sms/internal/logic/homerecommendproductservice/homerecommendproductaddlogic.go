package homerecommendproductservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// HomeRecommendProductAddLogic 人气商品推荐
/*
Author: LiuFeiHua
Date: 2024/5/14 9:33
*/
type HomeRecommendProductAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomeRecommendProductAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeRecommendProductAddLogic {
	return &HomeRecommendProductAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// HomeRecommendProductAdd 添加人气商品推荐
func (l *HomeRecommendProductAddLogic) HomeRecommendProductAdd(in *smsclient.HomeRecommendProductAddReq) (*smsclient.HomeRecommendProductAddResp, error) {
	q := query.SmsHomeRecommendProduct
	for _, data := range in.RecommendProductAddData {
		homeBrand, _ := q.WithContext(l.ctx).Where(q.ProductID.Eq(data.ProductId)).First()
		if homeBrand == nil {
			err := q.WithContext(l.ctx).Create(&model.SmsHomeRecommendProduct{
				ProductID:       data.ProductId,
				ProductName:     data.ProductName,
				RecommendStatus: data.RecommendStatus,
				Sort:            data.Sort,
			})

			if err != nil {
				return nil, err
			}
		}

	}

	return &smsclient.HomeRecommendProductAddResp{}, nil
}
