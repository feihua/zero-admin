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

// QuerySeckillSessionDetailLogic 查询秒杀场次详情
/*
Author: LiuFeiHua
Date: 2025/06/12 10:02:04
*/
type QuerySeckillSessionDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQuerySeckillSessionDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySeckillSessionDetailLogic {
	return &QuerySeckillSessionDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QuerySeckillSessionDetail 查询秒杀场次详情
func (l *QuerySeckillSessionDetailLogic) QuerySeckillSessionDetail(req *types.QuerySeckillSessionDetailReq) (resp *types.QuerySeckillSessionDetailResp, err error) {

	detail, err := l.svcCtx.SeckillSessionService.QuerySeckillSessionDetail(l.ctx, &smsclient.QuerySeckillSessionDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询秒杀场次详情失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	data := types.QuerySeckillSessionDetailData{
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

	}
	return &types.QuerySeckillSessionDetailResp{
		Code:    "000000",
		Message: "查询秒杀场次成功",
		Data:    data,
	}, nil
}
