package homerecommendproductservicelogic

import (
	"context"
	"zero-admin/rpc/model/smsmodel"
	"zero-admin/rpc/sms/smsclient"

	"zero-admin/rpc/sms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *HomeRecommendProductAddLogic) HomeRecommendProductAdd(in *smsclient.HomeRecommendProductAddReq) (*smsclient.HomeRecommendProductAddResp, error) {
	for _, data := range in.RecommendProductAddData {
		homeBrand, _ := l.svcCtx.SmsHomeBrandModel.FindOneByBrandId(l.ctx, data.ProductId)
		if homeBrand == nil {
			_, err := l.svcCtx.SmsHomeRecommendProductModel.Insert(l.ctx, &smsmodel.SmsHomeRecommendProduct{
				ProductId:       data.ProductId,
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
