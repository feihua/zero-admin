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

func (l *HomeAdvertiseDeleteLogic) HomeAdvertiseDelete(req types.DeleteHomeAdvertiseReq) (*types.DeleteHomeAdvertiseResp, error) {
	_, err := l.svcCtx.HomeAdvertiseService.HomeAdvertiseDelete(l.ctx, &smsclient.HomeAdvertiseDeleteReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据Id: %+v,删除首页广告异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("删除首页广告失败")
	}

	return &types.DeleteHomeAdvertiseResp{
		Code:    "000000",
		Message: "",
	}, nil
}
