package couponservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"time"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddCouponLogic 添加优惠券
/*
Author: LiuFeiHua
Date: 2024/6/12 17:34
*/
type AddCouponLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCouponLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCouponLogic {
	return &AddCouponLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddCoupon 添加优惠券
// 1.插入优惠券表
// 2.插入优惠券和商品关系表
// 3.插入优惠券和商品分类关系表
func (l *AddCouponLogic) AddCoupon(in *smsclient.AddOrUpdateCouponReq) (*smsclient.AddOrUpdateCouponResp, error) {
	StartTime, _ := time.Parse("2006-01-02 15:04:05", in.StartTime)
	EndTime, _ := time.Parse("2006-01-02 15:04:05", in.EndTime)
	EnableTime, _ := time.Parse("2006-01-02 15:04:05", in.EnableTime)

	// 1.插入优惠券表
	smsCoupon := &model.SmsCoupon{
		Type:         in.Type,         // 优惠券类型；0->全场赠券；1->会员赠券；2->购物赠券；3->注册赠券
		Name:         in.Name,         // 名称
		Platform:     in.Platform,     // 使用平台：0->全部；1->移动；2->PC
		Count:        in.Count,        // 数量
		Amount:       in.Amount,       // 金额
		PerLimit:     in.PerLimit,     // 每人限领张数
		MinPoint:     in.MinPoint,     // 使用门槛；0表示无门槛
		StartTime:    StartTime,       // 开始时间
		EndTime:      EndTime,         // 结束时间
		UseType:      in.UseType,      // 使用类型：0->全场通用；1->指定分类；2->指定商品
		Note:         in.Note,         // 备注
		PublishCount: in.PublishCount, // 发行数量
		UseCount:     in.UseCount,     // 已使用数量
		ReceiveCount: in.ReceiveCount, // 领取数量
		EnableTime:   EnableTime,      // 可以领取的日期
		Code:         in.Code,         // 优惠码
		MemberLevel:  in.MemberLevel,  // 可领取的会员类型：0->无限时
	}
	err := query.SmsCoupon.WithContext(l.ctx).Create(smsCoupon)

	if err != nil {
		return nil, err
	}

	// 2.插入优惠券和商品关系表
	if in.UseType == 2 {
		var list []*model.SmsCouponProductRelation
		for _, item := range in.CouponProductRelationList {
			list = append(list, &model.SmsCouponProductRelation{
				CouponID:    smsCoupon.ID,
				ProductID:   item.ProductId,
				ProductName: item.ProductName,
				ProductSn:   item.ProductSn,
			})
		}

		err = query.SmsCouponProductRelation.WithContext(l.ctx).CreateInBatches(list, len(list))
		if err != nil {
			return nil, err
		}
	}

	// 3.插入优惠券和商品分类关系表
	if in.UseType == 1 {
		var list []*model.SmsCouponProductCategoryRelation
		for _, item := range in.CouponProductCategoryRelationList {
			list = append(list, &model.SmsCouponProductCategoryRelation{
				CouponID:            smsCoupon.ID,
				ProductCategoryID:   item.ProductCategoryId,
				ProductCategoryName: item.ProductCategoryName,
				ParentCategoryName:  item.ParentCategoryName,
			})
		}

		err = query.SmsCouponProductCategoryRelation.WithContext(l.ctx).CreateInBatches(list, len(list))
		if err != nil {
			return nil, err
		}
	}

	return &smsclient.AddOrUpdateCouponResp{}, nil
}
