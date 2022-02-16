package logic

import (
	"context"
	"time"
	"zero-admin/rpc/model/omsmodel"
	"zero-admin/rpc/oms/internal/svc"
	"zero-admin/rpc/oms/oms"

	"github.com/zeromicro/go-zero/core/logx"
)

type CartItemAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCartItemAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartItemAddLogic {
	return &CartItemAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CartItemAddLogic) CartItemAdd(in *oms.CartItemAddReq) (*oms.CartItemAddResp, error) {
	createDate, _ := time.Parse("2006-01-02 15:04:05", in.CreateDate)
	modifyDate, _ := time.Parse("2006-01-02 15:04:05", in.ModifyDate)
	_, err := l.svcCtx.OmsCartItemModel.Insert(omsmodel.OmsCartItem{
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

	return &oms.CartItemAddResp{}, nil
}
