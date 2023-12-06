package cart

import (
	"context"
	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// CarItemtListPromotionLogic
/*
Author: LiuFeiHua
Date: 2023/12/6 16:18
*/
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

// CarItemtListPromotion 获取某个会员的购物车列表,包括促销信息
func (l *CarItemtListPromotionLogic) CarItemtListPromotion(req *types.CarItemListPromotionReq) (resp *types.CarItemtListPromotionResp, err error) {
	//1.获取会员购物车里面所有商品信息
	//memberId, _ := l.ctx.Value("memberId").(json.Number).Int64()
	//itemListResp, _ := l.svcCtx.CartItemService.CartItemList(l.ctx, &omsclient.CartItemListReq{MemberId: memberId})

	return
}
