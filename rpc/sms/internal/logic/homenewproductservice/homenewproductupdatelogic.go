package homenewproductservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeNewProductUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomeNewProductUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeNewProductUpdateLogic {
	return &HomeNewProductUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HomeNewProductUpdateLogic) HomeNewProductUpdate(in *smsclient.HomeNewProductUpdateReq) (*smsclient.HomeNewProductUpdateResp, error) {
	q := query.SmsHomeNewProduct
	_, err := q.WithContext(l.ctx).Updates(&model.SmsHomeNewProduct{
		ID:              in.Id,
		ProductID:       in.ProductId,
		ProductName:     in.ProductName,
		RecommendStatus: in.RecommendStatus,
		Sort:            in.Sort,
	})

	if err != nil {
		return nil, err
	}

	return &smsclient.HomeNewProductUpdateResp{}, nil
}
