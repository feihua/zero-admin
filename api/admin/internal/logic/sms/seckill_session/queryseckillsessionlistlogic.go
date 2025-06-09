package seckill_session

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

// QuerySeckillSessionListLogic 查询秒杀场次列表
/*
Author: LiuFeiHua
Date: 2025/06/12 10:02:04
*/
type QuerySeckillSessionListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQuerySeckillSessionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySeckillSessionListLogic {
	return &QuerySeckillSessionListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QuerySeckillSessionList 查询秒杀场次列表
func (l *QuerySeckillSessionListLogic) QuerySeckillSessionList(req *types.QuerySeckillSessionListReq) (resp *types.QuerySeckillSessionListResp, err error) {
	result, err := l.svcCtx.SeckillSessionService.QuerySeckillSessionList(l.ctx, &smsclient.QuerySeckillSessionListReq{
		PageNum:   req.Current,
		PageSize:  req.PageSize,
		Name:      req.Name,      // 场次名称
		StartTime: req.StartTime, // 开始时间
		EndTime:   req.EndTime,   // 结束时间
		Status:    req.Status,    // 状态：0-禁用，1-启用
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询字秒杀场次列表失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.QuerySeckillSessionListData

	for _, detail := range result.List {
		list = append(list, &types.QuerySeckillSessionListData{
			Id:         detail.Id,         // 秒杀场次ID
			Name:       detail.Name,       // 场次名称
			StartTime:  detail.StartTime,  // 开始时间
			EndTime:    detail.EndTime,    // 结束时间
			Status:     detail.Status,     // 状态：0-禁用，1-启用
			Sort:       detail.Sort,       // 排序
			CreateBy:   detail.CreateBy,   // 创建人ID
			CreateTime: detail.CreateTime, // 创建时间
			UpdateBy:   detail.UpdateBy,   // 更新人ID
			UpdateTime: detail.UpdateTime, // 更新时间

		})
	}

	return &types.QuerySeckillSessionListResp{
		Code:     "000000",
		Message:  "查询秒杀场次列表成功",
		Data:     list,
		Current:  req.Current,
		PageSize: req.PageSize,
		Total:    result.Total,
		Success:  true,
	}, nil
}
