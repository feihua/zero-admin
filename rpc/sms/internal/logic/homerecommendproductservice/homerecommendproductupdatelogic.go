package homerecommendproductservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/model/smsmodel"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeRecommendProductUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomeRecommendProductUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeRecommendProductUpdateLogic {
	return &HomeRecommendProductUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HomeRecommendProductUpdateLogic) HomeRecommendProductUpdate(in *smsclient.HomeRecommendProductUpdateReq) (*smsclient.HomeRecommendProductUpdateResp, error) {
	err := l.svcCtx.SmsHomeRecommendProductModel.Update(l.ctx, &smsmodel.SmsHomeRecommendProduct{
		Id:              in.Id,
		ProductId:       in.ProductId,
		ProductName:     in.ProductName,
		RecommendStatus: in.RecommendStatus,
		Sort:            in.Sort,
	})
	if err != nil {
		return nil, err
	}

	return &smsclient.HomeRecommendProductUpdateResp{}, nil
}
