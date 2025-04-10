package couponhistoryservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryCouponCountLogic 登录时获取用户还没有使用的获取优惠券数量
/*
Author: LiuFeiHua
Date: 2024/6/12 17:29
*/
type QueryCouponCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryCouponCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCouponCountLogic {
	return &QueryCouponCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryCouponCount 登录时获取用户还没有使用的获取优惠券数量
func (l *QueryCouponCountLogic) QueryCouponCount(in *smsclient.QueryCouponCountReq) (*smsclient.QueryCouponCountResp, error) {
	q := query.SmsCouponHistory
	count, err := q.WithContext(l.ctx).Where(q.MemberID.Eq(in.MemberId)).Count()

	if err != nil {
		logc.Errorf(l.ctx, "登录时获取用户还没有使用的获取优惠券数量失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("登录时获取用户还没有使用的获取优惠券数量失败")
	}
	return &smsclient.QueryCouponCountResp{
		Total: count,
	}, nil

}
