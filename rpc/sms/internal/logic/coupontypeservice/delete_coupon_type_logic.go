package coupontypeservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteCouponTypeLogic 删除优惠券类型
/*
Author: LiuFeiHua
Date: 2025/06/11 10:44:50
*/
type DeleteCouponTypeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCouponTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCouponTypeLogic {
	return &DeleteCouponTypeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteCouponType 删除优惠券类型
func (l *DeleteCouponTypeLogic) DeleteCouponType(in *smsclient.DeleteCouponTypeReq) (*smsclient.DeleteCouponTypeResp, error) {
	q := query.SmsCouponType

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除优惠券类型失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除优惠券类型失败")
	}

	return &smsclient.DeleteCouponTypeResp{}, nil
}
