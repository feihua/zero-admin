package home_advertise

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateHomeAdvertiseStatusLogic 更新首页轮播广告状态状态
/*
Author: LiuFeiHua
Date: 2025/06/12 10:40:15
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

// UpdateHomeAdvertiseStatus 更新首页轮播广告状态
func (l *UpdateHomeAdvertiseStatusLogic) UpdateHomeAdvertiseStatus(req *types.UpdateHomeAdvertiseStatusReq) (resp *types.BaseResp, err error) {

	_, err = l.svcCtx.HomeAdvertiseService.UpdateHomeAdvertiseStatus(l.ctx, &smsclient.UpdateHomeAdvertiseStatusReq{
		Ids:    req.Ids,    // 编号
		Status: req.Status, // 上下线状态：0->下线；1->上线

	})

	if err != nil {
		logc.Errorf(l.ctx, "更新首页轮播广告状态失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "更新首页轮播广告状态成功",
	}, nil
}
