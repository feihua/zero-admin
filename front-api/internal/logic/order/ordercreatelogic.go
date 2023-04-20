package order

import (
	"context"
	"encoding/json"

	"zero-admin/common/ctxdata"
	"zero-admin/common/errorx"
	"zero-admin/common/uniqueid"
	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"
	"zero-admin/rpc/oms/omsclient"
	"zero-admin/rpc/pms/pmsclient"
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

func (l *OrderCreateLogic) OrderCreate(req *types.OrderCreateReq) (resp *types.OrderCreateResp, err error) {
	goods := req.Goods
	l.Logger.Infof("goods", goods)
	paymenetId := req.PaymentId
	addressId := req.AddressId
	memberId := ctxdata.GetUidFromCtx(l.ctx)
	addressDetail := &umsclient.MemberReceiveAddressQueryDetailResp{
		Name:          "",
		PhoneNumber:   "",
		PostCode:      "",
		Province:      "",
		City:          "",
		Region:        "",
		DetailAddress: "",
	}
	var (
		resulterr  error
		isDelivery int64
	)

	//0自取，>0配送（选择配送地址）
	if addressId != 0 {
		isDelivery = 1
		addressDetail, resulterr = l.svcCtx.Ums.MemberReceiveAddressQueryDetail(l.ctx, &umsclient.MemberReceiveAddressQueryDetailReq{
			MemberId:  memberId,
			AddressID: addressId,
		})

		if resulterr != nil {
			return nil, resulterr
		}
	} else {
		isDelivery = 0
	}
	// 根据产品ID计算产品价格
	var omsgoods []*omsclient.Good
	totalAmount := 0.00

	for _, good := range goods {
		product, err := l.svcCtx.Pms.ProductDetailById(l.ctx, &pmsclient.ProductDetailByIdReq{Id: good.Id})
		if err != nil {
			return nil, err
		}
		omsgoods = append(omsgoods, &omsclient.Good{
			Id:  good.Id,
			Num: good.Num,
		})
		totalAmount += product.Price * float64(good.Num)
	}

	member, err := l.svcCtx.Ums.MemberInfo(l.ctx, &umsclient.MemberInfoReq{Id: memberId})
	if member == nil || err != nil {
		return nil, err
	}
	l.Logger.Infof("会员等级ID：%s", member.Member.LevelId)
	memberLevel, err := l.svcCtx.Ums.MemberLevelInfo(l.ctx, &umsclient.MemberLevelInfoReq{Id: member.Member.LevelId})
	if memberLevel == nil || err != nil {
		return nil, err
	}

	// 组装提交订单
	orderSn := uniqueid.GenSn(uniqueid.SN_PREFIX_HOMESTAY_ORDER)
	OrderAddResp, err := l.svcCtx.Oms.OrderAdd(l.ctx, &omsclient.OrderAddReq{
		MemberId:              memberId,
		CouponId:              0,
		OrderSn:               orderSn,
		CreateTime:            "",
		MemberUsername:        "",
		TotalAmount:           totalAmount,
		PayAmount:             totalAmount * memberLevel.Info.DiscountRate,
		FreightAmount:         0.00,
		PromotionAmount:       0.00,
		IntegrationAmount:     0.00,
		CouponAmount:          0.00,
		DiscountAmount:        0.00,
		PayType:               paymenetId,
		SourceType:            1,
		Status:                0,
		OrderType:             0,
		IsDelivery:            isDelivery,
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
		Goods:                 omsgoods,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("添加订单信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("添加订单信息失败")
	}

	return &types.OrderCreateResp{
		Errno: 0,
		Data: types.OrderData{
			OrderId: OrderAddResp.OrderId,
			OrderSn: OrderAddResp.OrderSn,
		},
		Errmsg: "下单完成",
	}, nil
}
