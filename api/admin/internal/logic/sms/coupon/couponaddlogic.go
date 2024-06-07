package coupon

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// CouponAddLogic 优惠券
/*
Author: LiuFeiHua
Date: 2024/5/15 10:30
*/
type CouponAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCouponAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) CouponAddLogic {
	return CouponAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CouponAdd 添加优惠券
func (l *CouponAddLogic) CouponAdd(req types.AddOrUpdateCouponReq) (*types.AddOrUpdateCouponResp, error) {
	var productList []*smsclient.CouponProductRelationListData
	for _, relation := range req.ProductRelationList {
		productList = append(productList, &smsclient.CouponProductRelationListData{
			CouponId:    relation.CouponId,
			ProductId:   relation.ProductId,
			ProductName: relation.ProductName,
			ProductSn:   relation.ProductSn,
		})
	}

	var categoryList []*smsclient.CouponProductCategoryRelationListData
	for _, relation := range req.ProductCategoryRelationList {
		categoryList = append(categoryList, &smsclient.CouponProductCategoryRelationListData{
			CouponId:            relation.CouponId,
			ProductCategoryId:   relation.ProductCategoryId,
			ProductCategoryName: relation.ProductCategoryName,
			ParentCategoryName:  relation.ParentCategoryName,
		})
	}

	_, err := l.svcCtx.CouponService.CouponAdd(l.ctx, &smsclient.CouponAddOrUpdateReq{
		Type:                              req.Type,
		Name:                              req.Name,
		Platform:                          req.Platform,
		Amount:                            req.Amount,
		PerLimit:                          req.PerLimit,
		MinPoint:                          req.MinPoint,
		StartTime:                         req.StartTime,
		EndTime:                           req.EndTime,
		UseType:                           req.UseType,
		Note:                              req.Note,
		PublishCount:                      req.PublishCount,
		EnableTime:                        req.EnableTime,
		Code:                              req.Code,
		MemberLevel:                       req.MemberLevel,
		CouponProductRelationList:         productList,
		CouponProductCategoryRelationList: categoryList,
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加优惠券信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("添加优惠券失败")
	}

	return &types.AddOrUpdateCouponResp{
		Code:    "000000",
		Message: "添加优惠券成功",
	}, nil
}
