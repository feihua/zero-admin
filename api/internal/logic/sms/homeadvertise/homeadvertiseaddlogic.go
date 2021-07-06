package logic

import (
	"context"
	"go-zero-admin/api/internal/common/errorx"
	"go-zero-admin/rpc/sms/smsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type HomeAdvertiseAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeAdvertiseAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomeAdvertiseAddLogic {
	return HomeAdvertiseAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeAdvertiseAddLogic) HomeAdvertiseAdd(req types.AddHomeAdvertiseReq) (*types.AddHomeAdvertiseResp, error) {
	_, err := l.svcCtx.Sms.HomeAdvertiseAdd(l.ctx, &smsclient.HomeAdvertiseAddReq{
		Name:       req.Name,
		Type:       req.Type,
		Pic:        req.Pic,
		StartTime:  req.StartTime,
		EndTime:    req.EndTime,
		Status:     req.Status,
		ClickCount: req.ClickCount,
		OrderCount: req.OrderCount,
		Url:        req.Url,
		Note:       req.Note,
		Sort:       req.Sort,
	})

	if err != nil {
		return nil, errorx.NewDefaultError("添加首页广告失败")
	}

	return &types.AddHomeAdvertiseResp{
		Code:    "000000",
		Message: "添加首页广告成功",
	}, nil
}
