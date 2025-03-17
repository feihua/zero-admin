package coupon

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

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
func (l *CouponAddLogic) CouponAdd(req *types.AddOrUpdateCouponReq) (*types.AddOrUpdateCouponResp, error) {
	var productList []*smsclient.CouponProductRelationData
	for _, relation := range req.ProductRelationList {
		productList = append(productList, &smsclient.CouponProductRelationData{
			CouponId:    relation.CouponId,
			ProductId:   relation.ProductId,
			ProductName: relation.ProductName,
			ProductSn:   relation.ProductSn,
		})
	}

	var categoryList []*smsclient.CouponProductCategoryRelationData
	for _, relation := range req.ProductCategoryRelationList {
		categoryList = append(categoryList, &smsclient.CouponProductCategoryRelationData{
			CouponId:            relation.CouponId,
			ProductCategoryId:   relation.ProductCategoryId,
			ProductCategoryName: relation.ProductCategoryName,
			ParentCategoryName:  relation.ParentCategoryName,
		})
	}

	_, err := l.svcCtx.CouponService.AddCoupon(l.ctx, &smsclient.AddOrUpdateCouponReq{
		Type:                              req.Type,         // 优惠券类型；0->全场赠券；1->会员赠券；2->购物赠券；3->注册赠券
		Name:                              req.Name,         // 名称
		Platform:                          req.Platform,     // 使用平台：0->全部；1->移动；2->PC
		Amount:                            req.Amount,       // 金额
		PerLimit:                          req.PerLimit,     // 每人限领张数
		MinPoint:                          req.MinPoint,     // 使用门槛；0表示无门槛
		StartTime:                         req.StartTime,    // 开始时间
		EndTime:                           req.EndTime,      // 结束时间
		UseType:                           req.UseType,      // 使用类型：0->全场通用；1->指定分类；2->指定商品
		Note:                              req.Note,         // 备注
		PublishCount:                      req.PublishCount, // 发行数量
		EnableTime:                        req.EnableTime,   // 可以领取的日期
		Code:                              req.Code,         // 优惠码
		MemberLevel:                       req.MemberLevel,  // 可领取的会员类型：0->无限时
		CouponProductRelationList:         productList,
		CouponProductCategoryRelationList: categoryList,
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加优惠券信息失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.AddOrUpdateCouponResp{
		Code:    "000000",
		Message: "添加优惠券成功",
	}, nil
}
