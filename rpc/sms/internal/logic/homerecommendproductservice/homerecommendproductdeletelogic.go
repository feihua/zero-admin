package homerecommendproductservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeRecommendProductDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomeRecommendProductDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeRecommendProductDeleteLogic {
	return &HomeRecommendProductDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HomeRecommendProductDeleteLogic) HomeRecommendProductDelete(in *smsclient.HomeRecommendProductDeleteReq) (*smsclient.HomeRecommendProductDeleteResp, error) {
	err := l.svcCtx.SmsHomeRecommendProductModel.DeleteByIds(l.ctx, in.Ids)

	if err != nil {
		return nil, err
	}

	return &smsclient.HomeRecommendProductDeleteResp{}, nil
}
