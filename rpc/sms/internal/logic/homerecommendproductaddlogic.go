package logic

import (
	"context"
	"zero-admin/rpc/model/smsmodel"

	"zero-admin/rpc/sms/internal/svc"
	"zero-admin/rpc/sms/sms"

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

func (l *HomeRecommendProductAddLogic) HomeRecommendProductAdd(in *sms.HomeRecommendProductAddReq) (*sms.HomeRecommendProductAddResp, error) {
	_, err := l.svcCtx.SmsHomeRecommendProductModel.Insert(smsmodel.SmsHomeRecommendProduct{
		ProductId:       in.ProductId,
		ProductName:     in.ProductName,
		RecommendStatus: in.RecommendStatus,
		Sort:            in.Sort,
	})
	if err != nil {
		return nil, err
	}

	return &sms.HomeRecommendProductAddResp{}, nil
}
