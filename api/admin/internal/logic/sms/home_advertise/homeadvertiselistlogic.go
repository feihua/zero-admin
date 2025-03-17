package home_advertise

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"
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
	resp, err := l.svcCtx.HomeAdvertiseService.QueryHomeAdvertiseList(l.ctx, &smsclient.QueryHomeAdvertiseListReq{
		Name:      strings.TrimSpace(req.Name),
		Type:      req.Type,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		Status:    req.Status,
		PageNum:   req.Current,
		PageSize:  req.PageSize,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询首页广告列表异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.ListHomeAdvertiseData

	for _, item := range resp.List {
		list = append(list, &types.ListHomeAdvertiseData{
			Id:         item.Id,         //
			Name:       item.Name,       // 名称
			Type:       item.Type,       // 轮播位置：0->PC首页轮播；1->app首页轮播
			Pic:        item.Pic,        // 图片地址
			StartTime:  item.StartTime,  // 开始时间
			EndTime:    item.EndTime,    // 结束时间
			Status:     item.Status,     // 上下线状态：0->下线；1->上线
			ClickCount: item.ClickCount, // 点击数
			OrderCount: item.OrderCount, // 下单数
			Url:        item.Url,        // 链接地址
			Note:       item.Note,       // 备注
			Sort:       item.Sort,       // 排序
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
