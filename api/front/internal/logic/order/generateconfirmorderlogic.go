package order

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/api/front/internal/logic/cart"
	"github.com/feihua/zero-admin/api/front/internal/logic/member/coupon"
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
	memberId, _ := l.ctx.Value("memberId").(json.Number).Int64()
	//1.获取购物车信息
	cartPromotionItemList := cart.QueryCartListPromotion(req.Ids, l.ctx, l.svcCtx)

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
			CreateDate:        item.CreateDate,
			ModifyDate:        item.ModifyDate,
			DeleteStatus:      item.DeleteStatus,
			ProductCategoryId: item.ProductCategoryId,
			ProductBrand:      item.ProductBrand,
			ProductSn:         item.ProductSn,
			ProductAttr:       item.ProductSn,
			PromotionMessage:  item.PromotionMessage,
			ReduceAmount:      item.ReduceAmount,
			RealStock:         item.RealStock,
			Integration:       item.Integration,
			Growth:            item.Growth,
		})
	}

	//2.获取用户收货地址列表
	addressListResp, _ := l.svcCtx.MemberReceiveAddressService.QueryMemberReceiveAddressList(l.ctx, &umsclient.QueryMemberReceiveAddressListReq{
		PageNum:  1,
		PageSize: 100,
		MemberId: memberId,
	})

	memberReceiveAddressList := make([]types.MemberReceiveAddressList, 0)
	for _, item := range addressListResp.List {
		memberReceiveAddressList = append(memberReceiveAddressList, types.MemberReceiveAddressList{
			Id:            item.Id,
			MemberId:      item.MemberId,
			Name:          item.MemberName,
			PhoneNumber:   item.PhoneNumber,
			DefaultStatus: item.DefaultStatus,
			PostCode:      item.PostCode,
			Province:      item.Province,
			City:          item.City,
			Region:        item.Region,
			DetailAddress: item.DetailAddress,
		})
	}
	//3.获取该用户所有未使用优惠券
	enableList, disableList := coupon.QueryCouponList(l.svcCtx, l.ctx, cartPromotionItemList)
	//4.获取用户积分
	memberInfo, _ := l.svcCtx.MemberService.QueryMemberDetail(l.ctx, &umsclient.QueryMemberDetailReq{
		Id: memberId,
	})
	//5.获取积分使用规则
	settingInfo, _ := l.svcCtx.IntegrationConsumeSettingService.QueryIntegrationConsumeSettingDetail(l.ctx, &umsclient.QueryIntegrationConsumeSettingDetailReq{
		Id: 1,
	})
	//6.计算总金额、活动优惠、应付金额
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
			MemberIntegration: int64(memberInfo.Integration),
			CalcAmount: types.CalcAmount{
				TotalAmount:     totalAmount,
				FreightAmount:   freightAmount,
				PromotionAmount: promotionAmount,
				PayAmount:       payAmount,
			},
		},
	}, nil
}
