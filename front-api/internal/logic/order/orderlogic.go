package order

import (
	"context"
	"encoding/json"
	"fmt"

	"zero-admin/common/ctxdata"
	"zero-admin/common/errorx"
	"zero-admin/common/uniqueid"
	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"
	"zero-admin/rpc/oms/omsclient"
	"zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderLogic {
	return &OrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderLogic) Order(req *types.OrderReq) (resp *types.OrderResp, err error) {
	goodsIds := req.GoodsIds
	// paymenetId := req.PaymentId
	// addressId := req.AddressId
	userId := ctxdata.GetUidFromCtx(l.ctx)
	totalAmount := 0.00
	for _, goodsId := range goodsIds {
		fmt.Println(goodsId)
		totalAmount += 20
	}
	orderSn := uniqueid.GenSn(uniqueid.SN_PREFIX_HOMESTAY_ORDER)
	OrderAddResp, err := l.svcCtx.Oms.OrderAdd(l.ctx, &omsclient.OrderAddReq{
		MemberId:              userId,
		CouponId:              0,
		OrderSn:               orderSn,
		CreateTime:            "",
		MemberUsername:        "",
		TotalAmount:           totalAmount,
		PayAmount:             totalAmount * 0.7,
		FreightAmount:         0.00,
		PromotionAmount:       0.00,
		IntegrationAmount:     0.00,
		CouponAmount:          0.00,
		DiscountAmount:        0.00,
		PayType:               0,
		SourceType:            1,
		Status:                0,
		OrderType:             0,
		DeliveryCompany:       "",
		DeliverySn:            "",
		AutoConfirmDay:        15,
		Integration:           0,
		Growth:                0,
		PromotionInfo:         "",
		BillType:              0,
		BillHeader:            "",
		BillContent:           "",
		BillReceiverPhone:     "",
		BillReceiverEmail:     "",
		ReceiverName:          "",
		ReceiverPhone:         "",
		ReceiverPostCode:      "",
		ReceiverProvince:      "",
		ReceiverCity:          "",
		ReceiverRegion:        "",
		ReceiverDetailAddress: "",
		Note:                  "",
		ConfirmStatus:         0,
		DeleteStatus:          0,
		UseIntegration:        0,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("添加订单信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("添加订单信息失败")
	}

	productList, err := l.svcCtx.Pms.ProductListByIds(l.ctx, &pmsclient.ProductListByIdsReq{
		Ids: goodsIds,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("查找订单商品有误")
	}

	for _, v := range productList.List {
		fmt.Println(v)
		_, err := l.svcCtx.Oms.OrderItemAdd(l.ctx, &omsclient.OrderItemAddReq{
			OrderId:           OrderAddResp.OrderId,
			OrderSn:           orderSn,
			ProductId:         v.Id,
			ProductPic:        v.Pic,
			ProductName:       v.Name,
			ProductBrand:      v.BrandName,
			ProductSn:         v.ProductSn,
			ProductPrice:      int64(v.Price),
			ProductQuantity:   0,
			ProductSkuId:      0,
			ProductSkuCode:    "",
			ProductCategoryId: v.ProductCategoryId,
			PromotionName:     "",
			PromotionAmount:   0.00,
			CouponAmount:      0.00,
			IntegrationAmount: 0.00,
			RealAmount:        0.00,
			GiftIntegration:   0,
			GiftGrowth:        v.GiftGrowth,
			ProductAttr:       "",
		})
		if err != nil {
			return nil, errorx.NewDefaultError("添加订单商品有误")
		}
	}

	fmt.Println(OrderAddResp)
	return &types.OrderResp{
		Errno:  0,
		Data:   "2020041715070001",
		Errmsg: "下单完成",
	}, nil
}
