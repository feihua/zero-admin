package homeadvertise

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"strings"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// HomeAdvertiseListLogic 首页轮播广告
/*
Author: LiuFeiHua
Date: 2024/5/13 17:33
*/
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

// HomeAdvertiseList 查询首页轮播广告
func (l *HomeAdvertiseListLogic) HomeAdvertiseList(req types.ListHomeAdvertiseReq) (*types.ListHomeAdvertiseResp, error) {
	resp, err := l.svcCtx.HomeAdvertiseService.HomeAdvertiseList(l.ctx, &smsclient.HomeAdvertiseListReq{
		Name:      strings.TrimSpace(req.Name),
		Type:      req.Type,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		Status:    req.Status,
		Current:   req.Current,
		PageSize:  req.PageSize,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询首页广告列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询首页广告失败")
	}

	var list []*types.ListHomeAdvertiseData

	for _, item := range resp.List {
		list = append(list, &types.ListHomeAdvertiseData{
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
