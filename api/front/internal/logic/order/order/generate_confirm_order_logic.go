package order

import (
	"context"
	"github.com/feihua/zero-admin/api/front/internal/logic/common"
	"github.com/feihua/zero-admin/api/front/internal/logic/member/coupon"
	"github.com/feihua/zero-admin/api/front/internal/logic/order/cart"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// GenerateConfirmOrderLogic
/*
Author: LiuFeiHua
Date: 2023/12/8 14:04
*/
type GenerateConfirmOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGenerateConfirmOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateConfirmOrderLogic {
	return &GenerateConfirmOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GenerateConfirmOrder 根据用户购物车信息生成确认单信息
// 1.获取购物车信息
// 2.获取用户收货地址列表
// 3.获取用户可用优惠券列表
// 4.获取用户积分
// 5.获取积分使用规则
// 6.计算总金额、活动优惠、应付金额
func (l *GenerateConfirmOrderLogic) GenerateConfirmOrder(req *types.GenerateConfirmOrderReq) (*types.GenerateConfirmOrderResp, error) {
	memberId, err := common.GetMemberId(l.ctx)
	if err != nil {
		return nil, err
	}
	// 1.获取购物车信息
	cartPromotionItemList, err := cart.QueryCartListPromotion(req.Ids, l.ctx, l.svcCtx)

	if err != nil {
		return nil, err
	}
	cartPromotionList := make([]types.CartPromotionItemList, 0)
	for _, item := range cartPromotionItemList {
		cartPromotionList = append(cartPromotionList, types.CartPromotionItemList{
			Id:                item.Id,
			ProductId:         item.ProductId,
			ProductSkuId:      item.ProductSkuId,
			MemberId:          item.MemberId,
			Quantity:          item.Quantity,
			Price:             item.Price,
			ProductPic:        item.ProductPic,
			ProductName:       item.ProductName,
			ProductSubTitle:   item.ProductSubTitle,
			ProductSkuCode:    item.ProductSkuCode,
			MemberNickname:    item.MemberNickname,
			CreateDate:        item.CreateTime,
			ModifyDate:        item.UpdateTime,
			DeleteStatus:      item.DeleteStatus,
			ProductCategoryId: item.ProductCategoryId,
			ProductBrand:      item.ProductBrand,
			ProductSn:         item.ProductSn,
			ProductAttr:       item.ProductAttr,
			PromotionMessage:  item.PromotionMessage,
			ReduceAmount:      item.ReduceAmount,
			RealStock:         item.RealStock,
			Integration:       item.Integration,
			Growth:            item.Growth,
		})
	}

	// 2.获取用户收货地址列表
	addressListResp, _ := l.svcCtx.MemberAddressService.QueryMemberAddressList(l.ctx, &umsclient.QueryMemberAddressListReq{
		PageNum:  1,
		PageSize: 100,
		MemberId: memberId,
	})

	memberReceiveAddressList := make([]types.MemberReceiveAddressList, 0)
	for _, detail := range addressListResp.List {
		memberReceiveAddressList = append(memberReceiveAddressList, types.MemberReceiveAddressList{
			Id:            detail.Id,            // 主键ID
			MemberId:      detail.MemberId,      // 会员ID
			ReceiverName:  detail.ReceiverName,  // 收货人姓名
			ReceiverPhone: detail.ReceiverPhone, // 收货人电话
			Province:      detail.Province,      // 省份
			City:          detail.City,          // 城市
			District:      detail.District,      // 区县
			DetailAddress: detail.DetailAddress, // 详细地址
			PostalCode:    detail.PostalCode,    // 邮政编码
			Tag:           detail.Tag,           // 地址标签：家、公司等
			IsDefault:     detail.IsDefault,     // 是否默认地址
		})
	}
	// 3.获取该用户所有未使用优惠券
	enableList, disableList := coupon.QueryCouponList(l.svcCtx, l.ctx, cartPromotionItemList)
	// 4.获取用户积分
	memberInfo, _ := l.svcCtx.MemberService.QueryMemberInfoDetail(l.ctx, &umsclient.QueryMemberInfoDetailReq{
		MemberId: memberId,
	})
	// 5.获取积分使用规则
	settingInfo, _ := l.svcCtx.MemberConsumeSettingService.QueryMemberConsumeSettingDetail(l.ctx, &umsclient.QueryMemberConsumeSettingDetailReq{
		Id: 1,
	})
	// 6.计算总金额、活动优惠、应付金额
	var totalAmount int64 = 0
	var freightAmount int64 = 0
	var promotionAmount int64 = 0
	var payAmount int64 = 0
	for _, item := range cartPromotionItemList {
		totalAmount = totalAmount + item.Price*int64(item.Quantity)
		promotionAmount = promotionAmount + item.ReduceAmount*int64(item.Quantity)
	}
	payAmount = totalAmount - promotionAmount
	return &types.GenerateConfirmOrderResp{
		Code:    0,
		Message: "操作成功",
		Data: types.OrderDetailModel{
			CartPromotionItemList:    cartPromotionList,
			MemberReceiveAddressList: memberReceiveAddressList,
			CouponHistoryDetailList: types.CouponListByCartData{
				EnableList:  enableList,
				DisableList: disableList,
			},
			IntegrationConsumeSetting: types.IntegrationConsumeSetting{
				Id:                 settingInfo.Id,
				DeductionPerAmount: settingInfo.DeductionPerAmount,
				MaxPercentPerOrder: settingInfo.MaxPercentPerOrder,
				UseUnit:            settingInfo.UseUnit,
				CouponStatus:       settingInfo.CouponStatus,
			},
			MemberIntegration: int64(memberInfo.Points),
			CalcAmount: types.CalcAmount{
				TotalAmount:     totalAmount,
				FreightAmount:   freightAmount,
				PromotionAmount: promotionAmount,
				PayAmount:       payAmount,
			},
		},
	}, nil
}
