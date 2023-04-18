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
	"zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderCreateLogic {
	return &OrderCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderCreateLogic) OrderCreate(req *types.OrderReq) (resp *types.OrderResp, err error) {
	goodsIds := req.GoodsIds
	// paymenetId := req.PaymentId
	addressId := req.AddressId
	memberId := ctxdata.GetUidFromCtx(l.ctx)

	addressDetail, err := l.svcCtx.Ums.MemberReceiveAddressQueryDetail(l.ctx, &umsclient.MemberReceiveAddressQueryDetailReq{
		MemberId:  memberId,
		AddressID: addressId,
	})

	if err != nil {
		return nil, err
	}

	totalAmount := 0.00
	for _, goodsId := range goodsIds {
		fmt.Println(goodsId)
		totalAmount += 20
	}
	orderSn := uniqueid.GenSn(uniqueid.SN_PREFIX_HOMESTAY_ORDER)

	OrderAddResp, err := l.svcCtx.Oms.OrderAdd(l.ctx, &omsclient.OrderAddReq{
		MemberId:              memberId,
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
		ReceiverName:          addressDetail.Name,
		ReceiverPhone:         addressDetail.PhoneNumber,
		ReceiverPostCode:      addressDetail.PostCode,
		ReceiverProvince:      addressDetail.Province,
		ReceiverCity:          addressDetail.City,
		ReceiverRegion:        addressDetail.Region,
		ReceiverDetailAddress: addressDetail.DetailAddress,
		Note:                  "",
		ConfirmStatus:         0,
		DeleteStatus:          0,
		UseIntegration:        0,
		Ids:                   goodsIds,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("添加订单信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("添加订单信息失败")
	}

	return &types.OrderResp{
		Errno: 0,
		Data: types.OrderData{
			OrderId: OrderAddResp.OrderId,
			OrderSn: OrderAddResp.OrderSn,
		},
		Errmsg: "下单完成",
	}, nil
}
