package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CouponHistoryDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCouponHistoryDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) CouponHistoryDeleteLogic {
	return CouponHistoryDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CouponHistoryDeleteLogic) CouponHistoryDelete(req types.DeleteCouponHistoryReq) (*types.DeleteCouponHistoryResp, error) {
	_, err := l.svcCtx.CouponHistoryService.CouponHistoryDelete(l.ctx, &smsclient.CouponHistoryDeleteReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据Id: %+v,删除优惠券使用记录异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("删除优惠券使用记录失败")
	}
	return &types.DeleteCouponHistoryResp{
		Code:    "000000",
		Message: "",
	}, nil
}
