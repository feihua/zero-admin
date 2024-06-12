package cartitemservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/oms/gen/model"
	"github.com/feihua/zero-admin/rpc/oms/gen/query"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateCartItemLogic 更新购物车表
/*
Author: LiuFeiHua
Date: 2024/6/12 10:05
*/
type UpdateCartItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCartItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCartItemLogic {
	return &UpdateCartItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateCartItem 更新购物车表
// 1.先更加id删除原来的
// 2.添加新的
func (l *UpdateCartItemLogic) UpdateCartItem(in *omsclient.UpdateCartItemReq) (*omsclient.UpdateCartItemResp, error) {
	q := query.OmsCartItem
	//1.先更加id删除原来的
	_, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.Id)).Delete()
	if err != nil {
		return nil, err
	}

	//2.添加新的
	err = q.WithContext(l.ctx).Create(&model.OmsCartItem{
		ProductID:         in.ProductId,
		ProductSkuID:      in.ProductSkuId,
		MemberID:          in.MemberId,
		Quantity:          in.Quantity,
		Price:             in.Price,
		ProductPic:        in.ProductPic,
		ProductName:       in.ProductName,
		ProductSubTitle:   in.ProductSubTitle,
		ProductSkuCode:    in.ProductSkuCode,
		MemberNickname:    in.MemberNickname,
		DeleteStatus:      in.DeleteStatus,
		ProductCategoryID: in.ProductCategoryId,
		ProductBrand:      in.ProductBrand,
		ProductSn:         in.ProductSn,
		ProductAttr:       in.ProductAttr,
	})

	if err != nil {
		return nil, err
	}

	return &omsclient.UpdateCartItemResp{}, nil
}
