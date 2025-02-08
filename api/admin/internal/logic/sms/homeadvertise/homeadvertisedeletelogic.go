package homeadvertise

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// HomeAdvertiseDeleteLogic 首页轮播广告
/*
Author: LiuFeiHua
Date: 2024/5/13 17:33
*/
type HomeAdvertiseDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeAdvertiseDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomeAdvertiseDeleteLogic {
	return HomeAdvertiseDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// HomeAdvertiseDelete 删除首页轮播广告
func (l *HomeAdvertiseDeleteLogic) HomeAdvertiseDelete(req types.DeleteHomeAdvertiseReq) (*types.DeleteHomeAdvertiseResp, error) {
	_, err := l.svcCtx.HomeAdvertiseService.DeleteHomeAdvertise(l.ctx, &smsclient.DeleteHomeAdvertiseReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据Id: %+v,删除首页广告异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.DeleteHomeAdvertiseResp{
		Code:    "000000",
		Message: "",
	}, nil
}
