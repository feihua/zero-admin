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

// CouponUpdateLogic 优惠券
/*
Author: LiuFeiHua
Date: 2024/5/15 10:39
*/
type CouponUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCouponUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponUpdateLogic {
	return &CouponUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CouponUpdate 更新优惠券
// 1.更新优惠券
// 2.插入优惠券和商品关系表
// 3.插入优惠券和商品分类关系表
func (l *CouponUpdateLogic) CouponUpdate(in *smsclient.CouponAddOrUpdateReq) (*smsclient.CouponAddOrUpdateResp, error) {
	StartTime, _ := time.Parse("2006-01-02 15:04:05", in.StartTime)
	EndTime, _ := time.Parse("2006-01-02 15:04:05", in.EndTime)
	EnableTime, _ := time.Parse("2006-01-02 15:04:05", in.EnableTime)

	//1.更新优惠券表
	_, err := query.SmsCoupon.WithContext(l.ctx).Updates(&model.SmsCoupon{
		ID:           in.Id,
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
		Code:         in.Code,
		MemberLevel:  in.MemberLevel,
	})

	if err != nil {
		return nil, err
	}

	// 2.插入优惠券和商品关系表
	if in.UseType == 2 {
		relation := query.SmsCouponProductRelation
		_, _ = relation.WithContext(l.ctx).Where(relation.CouponID.Eq(in.Id)).Delete()

		var list []*model.SmsCouponProductRelation
		for _, item := range in.CouponProductRelationList {
			list = append(list, &model.SmsCouponProductRelation{
				CouponID:    item.CouponId,
				ProductID:   item.ProductId,
				ProductName: item.ProductName,
				ProductSn:   item.ProductSn,
			})
		}

		err = relation.WithContext(l.ctx).CreateInBatches(list, len(list))
		if err != nil {
			return nil, err
		}
	}

	// 3.插入优惠券和商品分类关系表
	if in.UseType == 1 {
		relation := query.SmsCouponProductCategoryRelation
		_, _ = relation.WithContext(l.ctx).Where(relation.CouponID.Eq(in.Id)).Delete()

		var list []*model.SmsCouponProductCategoryRelation
		for _, item := range in.CouponProductCategoryRelationList {
			list = append(list, &model.SmsCouponProductCategoryRelation{
				CouponID:            item.CouponId,
				ProductCategoryID:   item.ProductCategoryId,
				ProductCategoryName: item.ProductCategoryName,
				ParentCategoryName:  item.ParentCategoryName,
			})
		}

		err = relation.WithContext(l.ctx).CreateInBatches(list, len(list))
		if err != nil {
			return nil, err
		}
	}

	return &smsclient.CouponAddOrUpdateResp{}, nil
}
