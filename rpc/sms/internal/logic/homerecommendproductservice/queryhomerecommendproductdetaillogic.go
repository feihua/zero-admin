package homerecommendproductservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryHomeRecommendProductDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryHomeRecommendProductDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryHomeRecommendProductDetailLogic {
	return &QueryHomeRecommendProductDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询人气推荐商品表详情
func (l *QueryHomeRecommendProductDetailLogic) QueryHomeRecommendProductDetail(in *smsclient.QueryHomeRecommendProductDetailReq) (*smsclient.QueryHomeRecommendProductDetailResp, error) {
	// todo: add your logic here and delete this line

	return &smsclient.QueryHomeRecommendProductDetailResp{}, nil
}
