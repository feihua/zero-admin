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

// DeleteHomeAdvertiseLogic 删除首页轮播广告
/*
Author: LiuFeiHua
Date: 2025/06/12 10:40:15
*/
type DeleteHomeAdvertiseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteHomeAdvertiseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteHomeAdvertiseLogic {
	return &DeleteHomeAdvertiseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteHomeAdvertise 删除首页轮播广告
func (l *DeleteHomeAdvertiseLogic) DeleteHomeAdvertise(req *types.DeleteHomeAdvertiseReq) (resp *types.BaseResp, err error) {
	_, err = l.svcCtx.HomeAdvertiseService.DeleteHomeAdvertise(l.ctx, &smsclient.DeleteHomeAdvertiseReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "删除首页轮播广告失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "删除首页轮播广告成功",
	}, nil
}
