package logic

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/api/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/feihua/zero-admin/api/internal/svc"
	"github.com/feihua/zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeAdvertiseListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeAdvertiseListLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomeAdvertiseListLogic {
	return HomeAdvertiseListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeAdvertiseListLogic) HomeAdvertiseList(req types.ListHomeAdvertiseReq) (*types.ListHomeAdvertiseResp, error) {
	resp, err := l.svcCtx.HomeAdvertiseService.HomeAdvertiseList(l.ctx, &smsclient.HomeAdvertiseListReq{
		Name:      req.Name,
		Type:      req.Type,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		Status:    req.Status,
		Current:   req.Current,
		PageSize:  req.PageSize,
	})

	if err != nil {
		data, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("参数: %s,查询首页广告列表异常:%s", string(data), err.Error())
		return nil, errorx.NewDefaultError("查询首页广告失败")
	}

	var list []*types.ListtHomeAdvertiseData

	for _, item := range resp.List {
		list = append(list, &types.ListtHomeAdvertiseData{
			Id:         item.Id,
			Name:       item.Name,
			Type:       item.Type,
			Pic:        item.Pic,
			StartTime:  item.StartTime,
			EndTime:    item.EndTime,
			Status:     item.Status,
			ClickCount: item.ClickCount,
			OrderCount: item.OrderCount,
			Url:        item.Url,
			Note:       item.Note,
			Sort:       item.Sort,
		})
	}

	return &types.ListHomeAdvertiseResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
		Code:     "000000",
		Message:  "查询首页广告成功",
	}, nil
}
