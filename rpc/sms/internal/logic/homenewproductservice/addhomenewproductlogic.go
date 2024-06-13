package homenewproductservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddHomeNewProductLogic 添加新鲜好物表
/*
Author: LiuFeiHua
Date: 2024/6/12 17:54
*/
type AddHomeNewProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddHomeNewProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddHomeNewProductLogic {
	return &AddHomeNewProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddHomeNewProduct 添加新鲜好物表
func (l *AddHomeNewProductLogic) AddHomeNewProduct(in *smsclient.AddHomeNewProductReq) (*smsclient.AddHomeNewProductResp, error) {
	for _, data := range in.NewProductAddData {
		q := query.SmsHomeNewProduct
		homeBrand, _ := q.WithContext(l.ctx).Where(q.ProductID.Eq(data.ProductId)).First()
		if homeBrand == nil {
			err := q.WithContext(l.ctx).Create(&model.SmsHomeNewProduct{
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

	return &smsclient.AddHomeNewProductResp{}, nil
}
