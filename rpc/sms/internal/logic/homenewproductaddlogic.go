package logic

import (
	"context"
	"go-zero-admin/rpc/model/smsmodel"

	"go-zero-admin/rpc/sms/internal/svc"
	"go-zero-admin/rpc/sms/sms"

	"github.com/tal-tech/go-zero/core/logx"
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
	_, err := l.svcCtx.SmsHomeNewProductModel.Insert(smsmodel.SmsHomeNewProduct{
		ProductId:       in.ProductId,
		ProductName:     in.ProductName,
		RecommendStatus: in.RecommendStatus,
		Sort:            in.Sort,
	})
	if err != nil {
		return nil, err
	}

	return &sms.HomeNewProductAddResp{}, nil
}
