package homerecommendproductservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddHomeRecommendProductLogic 添加人气推荐商品
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

// AddHomeRecommendProduct 添加人气推荐商品
func (l *AddHomeRecommendProductLogic) AddHomeRecommendProduct(in *smsclient.AddHomeRecommendProductReq) (*smsclient.AddHomeRecommendProductResp, error) {
	q := query.SmsHomeRecommendProduct
	for _, data := range in.RecommendProductAddData {
		homeBrand, _ := q.WithContext(l.ctx).Where(q.ProductID.Eq(data.ProductId)).First()
		if homeBrand == nil {
			err := q.WithContext(l.ctx).Create(&model.SmsHomeRecommendProduct{
				ProductID:       data.ProductId,       // 商品id
				ProductName:     data.ProductName,     // 商品名称
				RecommendStatus: data.RecommendStatus, // 推荐状态：0->不推荐;1->推荐
				Sort:            data.Sort,            // 排序
			})

			if err != nil {
				logc.Errorf(l.ctx, "添加人气推荐商品失败,参数:%+v,异常:%s", in, err.Error())
				return nil, errors.New("添加人气推荐商品失败")
			}
		}

	}

	return &smsclient.AddHomeRecommendProductResp{}, nil
}
