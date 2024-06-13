package homerecommendproductservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddHomeRecommendProductLogic 添加人气推荐商品表
/*
Author: LiuFeiHua
Date: 2024/6/12 17:56
*/
type AddHomeRecommendProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddHomeRecommendProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddHomeRecommendProductLogic {
	return &AddHomeRecommendProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddHomeRecommendProduct 添加人气推荐商品表
func (l *AddHomeRecommendProductLogic) AddHomeRecommendProduct(in *smsclient.AddHomeRecommendProductReq) (*smsclient.AddHomeRecommendProductResp, error) {
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

	return &smsclient.AddHomeRecommendProductResp{}, nil
}
