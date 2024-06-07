package couponservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"time"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

// CouponAddLogic
/*
Author: LiuFeiHua
Date: 2024/5/15 10:31
*/
type CouponAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCouponAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponAddLogic {
	return &CouponAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CouponAdd 添加优惠券
// 1.插入优惠券表
// 2.插入优惠券和商品关系表
// 3.插入优惠券和商品分类关系表
func (l *CouponAddLogic) CouponAdd(in *smsclient.CouponAddOrUpdateReq) (*smsclient.CouponAddOrUpdateResp, error) {
	StartTime, _ := time.Parse("2006-01-02 15:04:05", in.StartTime)
	EndTime, _ := time.Parse("2006-01-02 15:04:05", in.EndTime)
	EnableTime, _ := time.Parse("2006-01-02 15:04:05", in.EnableTime)

	//1.插入优惠券表
	smsCoupon := &model.SmsCoupon{
		Type:         in.Type,
		Name:         in.Name,
		Platform:     in.Platform,
		Amount:       in.Amount,
		PerLimit:     in.PerLimit,
		MinPoint:     in.MinPoint,
		StartTime:    StartTime,
		EndTime:      EndTime,
		UseType:      in.UseType,
		Note:         in.Note,
		PublishCount: in.PublishCount,
		UseCount:     in.UseCount,
		ReceiveCount: in.ReceiveCount,
		EnableTime:   EnableTime,
		Code:         "todo",
		MemberLevel:  in.MemberLevel,
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

	return &smsclient.CouponAddOrUpdateResp{}, nil
}
