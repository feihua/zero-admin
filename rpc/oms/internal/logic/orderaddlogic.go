package logic

import (
	"context"
	"time"
	"zero-admin/common/errorx"
	"zero-admin/rpc/model/omsmodel"
	"zero-admin/rpc/model/pmsmodel"
	"zero-admin/rpc/oms/internal/svc"
	"zero-admin/rpc/oms/oms"
	"zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type OrderAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderAddLogic {
	return &OrderAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrderAddLogic) OrderAdd(in *oms.OrderAddReq) (*oms.OrderAddResp, error) {

	createTime, _ := time.Parse("2006-01-02 15:04:05", in.CreateTime)
	paymentTime, _ := time.Parse("2006-01-02 15:04:05", in.PaymentTime)
	deliveryTime, _ := time.Parse("2006-01-02 15:04:05", in.DeliveryTime)
	receiveTime, _ := time.Parse("2006-01-02 15:04:05", in.ReceiveTime)
	commentTime, _ := time.Parse("2006-01-02 15:04:05", in.CommentTime)
	modifyTime, _ := time.Parse("2006-01-02 15:04:05", in.ModifyTime)

	var lastInsertId int64
	if err := l.svcCtx.OmsOrderModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		result, err := l.svcCtx.OmsOrderModel.InsertTx(l.ctx, session, &omsmodel.OmsOrder{
			MemberId:              in.MemberId,
			CouponId:              in.CouponId,
			OrderSn:               in.OrderSn,
			CreateTime:            createTime,
			MemberUsername:        in.MemberUsername,
			TotalAmount:           in.TotalAmount,
			PayAmount:             in.PayAmount,
			FreightAmount:         in.FreightAmount,
			PromotionAmount:       in.PromotionAmount,
			IntegrationAmount:     in.IntegrationAmount,
			CouponAmount:          in.CouponAmount,
			DiscountAmount:        in.DiscountAmount,
			PayType:               in.PayType,
			SourceType:            in.SourceType,
			Status:                in.Status,
			OrderType:             in.OrderType,
			IsDelivery:            in.IsDelivery,
			DeliveryCompany:       in.DeliveryCompany,
			DeliverySn:            in.DeliverySn,
			AutoConfirmDay:        in.AutoConfirmDay,
			Integration:           in.Integration,
			Growth:                in.Growth,
			PromotionInfo:         in.PromotionInfo,
			BillType:              in.BillType,
			BillHeader:            in.BillHeader,
			BillContent:           in.BillContent,
			BillReceiverPhone:     in.BillReceiverPhone,
			BillReceiverEmail:     in.BillReceiverEmail,
			ReceiverName:          in.ReceiverName,
			ReceiverPhone:         in.ReceiverPhone,
			ReceiverPostCode:      in.ReceiverPostCode,
			ReceiverProvince:      in.ReceiverProvince,
			ReceiverCity:          in.ReceiverCity,
			ReceiverRegion:        in.ReceiverRegion,
			ReceiverDetailAddress: in.ReceiverDetailAddress,
			Note:                  in.Note,
			ConfirmStatus:         in.ConfirmStatus,
			DeleteStatus:          in.DeleteStatus,
			UseIntegration:        in.UseIntegration,
			PaymentTime:           paymentTime,
			DeliveryTime:          deliveryTime,
			ReceiveTime:           receiveTime,
			CommentTime:           commentTime,
			ModifyTime:            modifyTime,
		})
		if err != nil {
			return err
		}
		lastInsertId, err = result.LastInsertId()
		if err != nil {
			return err
		}
		var goodsIds []int64
		for _, good := range in.Goods {
			goodsIds = append(goodsIds, good.Id)
		}
		var productList *[]pmsmodel.PmsProduct
		err = l.svcCtx.DbEngin.Find(&productList, goodsIds).Error
		// productList, err := l.svcCtx.PmsProductModel.ProductListByIds(in.Ids)
		if err != nil {
			logx.WithContext(l.ctx).Infof(err.Error())
			return err
		}
		for k, v := range *productList {
			_, err := l.svcCtx.OmsOrderItemModel.InsertTx(l.ctx, session, &omsmodel.OmsOrderItem{
				OrderId:           lastInsertId,
				OrderSn:           in.OrderSn,
				ProductId:         v.Id,
				ProductPic:        v.Pic,
				ProductName:       v.Name,
				ProductBrand:      v.BrandName,
				ProductSn:         v.ProductSn,
				ProductPrice:      float64(v.Price),
				ProductQuantity:   in.Goods[k].Num,
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
				return err
			}
		}

		return nil
	}); err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}

	return &omsclient.OrderAddResp{
		OrderId: lastInsertId,
		OrderSn: in.OrderSn,
	}, nil

}
