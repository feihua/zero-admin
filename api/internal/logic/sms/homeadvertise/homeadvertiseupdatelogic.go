package logic

import (
	"context"
	"encoding/json"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/sms/smsclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeAdvertiseUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeAdvertiseUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomeAdvertiseUpdateLogic {
	return HomeAdvertiseUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeAdvertiseUpdateLogic) HomeAdvertiseUpdate(req types.UpdateHomeAdvertiseReq) (*types.UpdateHomeAdvertiseResp, error) {
	_, err := l.svcCtx.Sms.HomeAdvertiseUpdate(l.ctx, &smsclient.HomeAdvertiseUpdateReq{
		Id:         req.Id,
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
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("更新首页广告信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("更新首页广告失败")
	}

	return &types.UpdateHomeAdvertiseResp{
		Code:    "000000",
		Message: "更新首页广告成功",
	}, nil
}
