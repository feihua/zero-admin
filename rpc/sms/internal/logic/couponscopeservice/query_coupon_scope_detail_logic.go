package couponscopeservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

// QueryCouponScopeDetailLogic 查询优惠券使用范围详情
/*
Author: LiuFeiHua
Date: 2025/06/11 10:44:50
*/
type QueryCouponScopeDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryCouponScopeDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCouponScopeDetailLogic {
	return &QueryCouponScopeDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryCouponScopeDetail 查询优惠券使用范围详情
func (l *QueryCouponScopeDetailLogic) QueryCouponScopeDetail(in *smsclient.QueryCouponScopeDetailReq) (*smsclient.QueryCouponScopeDetailResp, error) {
	item, err := query.SmsCouponScope.WithContext(l.ctx).Where(query.SmsCouponScope.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "优惠券使用范围不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("优惠券使用范围不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询优惠券使用范围异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询优惠券使用范围异常")
	}

	data := &smsclient.QueryCouponScopeDetailResp{
		Id:         item.ID,                                       //
		CouponId:   item.CouponID,                                 // 优惠券ID
		ScopeType:  item.ScopeType,                                // 范围类型：0-全场，1-分类，2-商品
		ScopeId:    item.ScopeID,                                  // 范围ID（分类ID或商品ID）
		CreateTime: item.CreateTime.Format("2006-01-02 15:04:05"), // 创建时间
	}

	return data, nil
}
