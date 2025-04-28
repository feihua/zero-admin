package home_advertise

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/common/res"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateHomeAdvertiseStatusLogic 首页轮播广告
/*
Author: LiuFeiHua
Date: 2024/5/13 17:33
*/
type UpdateHomeAdvertiseStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateHomeAdvertiseStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateHomeAdvertiseStatusLogic {
	return &UpdateHomeAdvertiseStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateHomeAdvertiseStatus 修改上下线状态
func (l *UpdateHomeAdvertiseStatusLogic) UpdateHomeAdvertiseStatus(req *types.UpdateHomeAdvertiseStatusReq) (resp *types.BaseResp, err error) {
	_, err = l.svcCtx.HomeAdvertiseService.UpdateHomeAdvertiseStatus(l.ctx, &smsclient.UpdateHomeAdvertiseStatusReq{
		Ids:    req.Ids,
		Status: req.Status,
	})

	if err != nil {
		logc.Errorf(l.ctx, "修改上下线状态失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return res.Success()
}
