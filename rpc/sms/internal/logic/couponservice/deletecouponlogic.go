package couponservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteCouponLogic 删除优惠券
/*
Author: LiuFeiHua
Date: 2024/6/12 17:34
*/
type DeleteCouponLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCouponLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCouponLogic {
	return &DeleteCouponLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteCoupon 删除优惠券
// 1.删除优惠券
// 2.删除商品关联
// 3.删除商品分类关联
func (l *DeleteCouponLogic) DeleteCoupon(in *smsclient.DeleteCouponReq) (*smsclient.DeleteCouponResp, error) {
	// 1.删除优惠券
	q := query.SmsCoupon
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除优惠券失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除优惠券失败")
	}

	// 2.删除商品关联
	relation := query.SmsCouponProductRelation
	_, err = relation.WithContext(l.ctx).Where(relation.CouponID.In(in.Ids...)).Delete()
	if err != nil {
		logc.Errorf(l.ctx, "删除优惠券失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除优惠券失败")
	}
	// 3.删除商品分类关联
	categoryRelation := query.SmsCouponProductCategoryRelation
	_, err = categoryRelation.WithContext(l.ctx).Where(categoryRelation.CouponID.In(in.Ids...)).Delete()
	if err != nil {
		logc.Errorf(l.ctx, "删除优惠券失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除优惠券失败")
	}

	return &smsclient.DeleteCouponResp{}, nil
}
