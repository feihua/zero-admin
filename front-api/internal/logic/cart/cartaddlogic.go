package cart

import (
	"context"
	"zero-admin/rpc/oms/omsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCartAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) CartAddLogic {
	return CartAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CartAddLogic) CartAdd(req types.CartAddReq) (resp *types.CartAddResp, err error) {
	l.svcCtx.Oms.CartItemAdd(l.ctx, &omsclient.CartItemAddReq{
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

	return &types.CartAddResp{
		Errno:  0,
		Errmsg: "",
	}, nil
}
