package logic

import (
	"context"
	"time"
	"zero-admin/rpc/model/omsmodel"

	"zero-admin/rpc/oms/internal/svc"
	"zero-admin/rpc/oms/oms"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartItemUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCartItemUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartItemUpdateLogic {
	return &CartItemUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CartItemUpdateLogic) CartItemUpdate(in *oms.CartItemUpdateReq) (*oms.CartItemUpdateResp, error) {
	createDate, _ := time.Parse("2006-01-02 15:04:05", in.CreateDate)
	modifyDate, _ := time.Parse("2006-01-02 15:04:05", in.ModifyDate)
	err := l.svcCtx.OmsCartItemModel.Update(omsmodel.OmsCartItem{
		Id:                in.Id,
		ProductId:         in.ProductId,
		ProductSkuId:      in.ProductSkuId,
		MemberId:          in.MemberId,
		Quantity:          in.Quantity,
		Price:             float64(in.Price),
		ProductPic:        in.ProductPic,
		ProductName:       in.ProductName,
		ProductSubTitle:   in.ProductSubTitle,
		ProductSkuCode:    in.ProductSkuCode,
		MemberNickname:    in.MemberNickname,
		CreateDate:        createDate,
		ModifyDate:        modifyDate,
		DeleteStatus:      in.DeleteStatus,
		ProductCategoryId: in.ProductCategoryId,
		ProductBrand:      in.ProductBrand,
		ProductSn:         in.ProductSn,
		ProductAttr:       in.ProductAttr,
	})
	if err != nil {
		return nil, err
	}

	return &oms.CartItemUpdateResp{}, nil
}
