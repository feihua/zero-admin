package logic

import (
	"context"
	"zero-admin/rpc/model/smsmodel"

	"zero-admin/rpc/sms/internal/svc"
	"zero-admin/rpc/sms/sms"

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

func (l *HomeNewProductAddLogic) HomeNewProductAdd(in *sms.HomeNewProductAddReq) (*sms.HomeNewProductAddResp, error) {
	for _, data := range in.NewProductAddData {
		homeBrand, _ := l.svcCtx.SmsHomeBrandModel.FindOneByBrandId(l.ctx, data.ProductId)
		if homeBrand == nil {
			_, err := l.svcCtx.SmsHomeNewProductModel.Insert(l.ctx, &smsmodel.SmsHomeNewProduct{
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

	return &sms.HomeNewProductAddResp{}, nil
}
