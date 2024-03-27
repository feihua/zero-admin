package flashpromotionproductrelationservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryFlashPromotionByProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryFlashPromotionByProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryFlashPromotionByProductLogic {
	return &QueryFlashPromotionByProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryFlashPromotionByProduct 根据商品id查询
func (l *QueryFlashPromotionByProductLogic) QueryFlashPromotionByProduct(in *smsclient.QueryFlashPromotionByProductReq) (*smsclient.QueryFlashPromotionByProductResp, error) {
	// todo: add your logic here and delete this line

	return &smsclient.QueryFlashPromotionByProductResp{}, nil
}
