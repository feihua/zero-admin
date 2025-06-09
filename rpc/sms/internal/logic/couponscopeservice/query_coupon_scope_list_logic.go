package couponscopeservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// QueryCouponScopeListLogic 查询优惠券使用范围列表
/*
Author: LiuFeiHua
Date: 2025/06/11 10:44:50
*/
type QueryCouponScopeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryCouponScopeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCouponScopeListLogic {
	return &QueryCouponScopeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryCouponScopeList 查询优惠券使用范围列表
func (l *QueryCouponScopeListLogic) QueryCouponScopeList(in *smsclient.QueryCouponScopeListReq) (*smsclient.QueryCouponScopeListResp, error) {
	couponScope := query.SmsCouponScope
	q := couponScope.WithContext(l.ctx)
	if in.CouponId != 0 {
		q = q.Where(couponScope.CouponID.Eq(in.CouponId))
	}
	if in.ScopeType != 3 {
		q = q.Where(couponScope.ScopeType.Eq(in.ScopeType))
	}
	if in.ScopeId != 0 {
		q = q.Where(couponScope.ScopeID.Eq(in.ScopeId))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询优惠券使用范围列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询优惠券使用范围列表失败")
	}

	var list []*smsclient.CouponScopeListData

	for _, item := range result {
		list = append(list, &smsclient.CouponScopeListData{
			Id:         item.ID,                                       //
			CouponId:   item.CouponID,                                 // 优惠券ID
			ScopeType:  item.ScopeType,                                // 范围类型：0-全场，1-分类，2-商品
			ScopeId:    item.ScopeID,                                  // 范围ID（分类ID或商品ID）
			CreateTime: item.CreateTime.Format("2006-01-02 15:04:05"), // 创建时间

		})
	}

	return &smsclient.QueryCouponScopeListResp{
		Total: count,
		List:  list,
	}, nil
}
