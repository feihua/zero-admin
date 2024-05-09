package homenewproductservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeNewProductAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomeNewProductAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeNewProductAddLogic {
	return &HomeNewProductAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HomeNewProductAddLogic) HomeNewProductAdd(in *smsclient.HomeNewProductAddReq) (*smsclient.HomeNewProductAddResp, error) {
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

	return &smsclient.HomeNewProductAddResp{}, nil
}
