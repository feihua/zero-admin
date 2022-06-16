package cart

import (
	"context"
	"zero-admin/rpc/oms/omsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCartUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) CartUpdateLogic {
	return CartUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CartUpdateLogic) CartUpdate(req types.CartUpdateReq) (resp *types.CartUpdateResp, err error) {
	l.svcCtx.Oms.CartItemUpdate(l.ctx, &omsclient.CartItemUpdateReq{
		ProductId:         0,
		ProductSkuId:      0,
		MemberId:          0,
		Quantity:          0,
		Price:             0,
		ProductPic:        "",
		ProductName:       "",
		ProductSubTitle:   "",
		ProductSkuCode:    "",
		MemberNickname:    "",
		CreateDate:        "",
		ModifyDate:        "",
		DeleteStatus:      0,
		ProductCategoryId: 0,
		ProductBrand:      "",
		ProductSn:         "",
		ProductAttr:       "",
	})

	return &types.CartUpdateResp{
		Errno:  0,
		Errmsg: "",
	}, nil
}
