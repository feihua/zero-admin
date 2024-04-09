package cart

import (
	"context"

	"github.com/feihua/zero-admin/api/web/internal/svc"
	"github.com/feihua/zero-admin/api/web/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CarItemtListPromotionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCarItemtListPromotionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CarItemtListPromotionLogic {
	return &CarItemtListPromotionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CarItemtListPromotionLogic) CarItemtListPromotion(req *types.CarItemListPromotionReq) (resp *types.CarItemtListPromotionResp, err error) {
	// todo: add your logic here and delete this line

	return
}
