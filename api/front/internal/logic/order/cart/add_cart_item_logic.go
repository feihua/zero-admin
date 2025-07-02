package cart

import (
	"context"
	"github.com/feihua/zero-admin/api/front/internal/logic/common"
	"github.com/feihua/zero-admin/pkg/errorx"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCartItemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddCartItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCartItemLogic {
	return &AddCartItemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddCartItem 添加商品进购物车
func (l *AddCartItemLogic) AddCartItem(req *types.CartItemReq) (resp *types.CartItemResp, err error) {
	memberId, err := common.GetMemberId(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.CartItemService.AddCartItem(l.ctx, &omsclient.AddCartItemReq{
		MemberId:          memberId,              // 会员ID
		ProductId:         req.ProductId,         // 商品ID
		ProductSkuId:      req.ProductSkuId,      // 商品SKU ID
		Quantity:          req.Quantity,          // 购买数量
		Price:             req.Price,             // 添加到购物车时的价格
		Selected:          req.Selected,          // 是否选中 0-未选中 1-选中
		ProductName:       req.ProductName,       // 商品名称
		ProductSubTitle:   req.ProductSubTitle,   // 商品副标题
		ProductPic:        req.ProductPic,        // 商品主图URL
		ProductSkuCode:    req.ProductSkuCode,    // 商品SKU编码
		ProductSn:         req.ProductSn,         // 商品货号
		ProductBrand:      req.ProductBrand,      // 商品品牌
		ProductCategoryId: req.ProductCategoryId, // 商品分类ID
		ProductAttr:       req.ProductAttr,       // 商品销售属性JSON
		MemberNickname:    req.MemberNickname,    // 会员昵称
		Source:            req.Source,            // 来源 1-PC 2-H5 3-小程序 4-APP
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加购物车失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.CartItemResp{
		Code:    0,
		Message: "添加购物车成功",
	}, nil
}
