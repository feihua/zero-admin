package seckill_activity

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

// QuerySeckillActivityDetailLogic 查询秒杀活动详情
/*
Author: LiuFeiHua
Date: 2025/06/12 10:02:03
*/
type QuerySeckillActivityDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQuerySeckillActivityDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySeckillActivityDetailLogic {
	return &QuerySeckillActivityDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QuerySeckillActivityDetail 查询秒杀活动详情
func (l *QuerySeckillActivityDetailLogic) QuerySeckillActivityDetail(req *types.QuerySeckillActivityDetailReq) (resp *types.QuerySeckillActivityDetailResp, err error) {

	detail, err := l.svcCtx.SeckillActivityService.QuerySeckillActivityDetail(l.ctx, &smsclient.QuerySeckillActivityDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询秒杀活动详情失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	data := types.QuerySeckillActivityDetailData{
		Id:          detail.Id,          // 编号
		Name:        detail.Name,        // 活动名称
		Description: detail.Description, // 活动描述
		StartTime:   detail.StartTime,   // 开始时间
		EndTime:     detail.EndTime,     // 结束时间
		Status:      detail.Status,      // 活动状态：0-未开始，1-进行中，2-已结束，3-已取消
		IsEnabled:   detail.IsEnabled,   // 是否启用
		CreateBy:    detail.CreateBy,    // 创建人ID
		CreateTime:  detail.CreateTime,  // 创建时间
		UpdateBy:    detail.UpdateBy,    // 更新人ID
		UpdateTime:  detail.UpdateTime,  // 更新时间

	}
	return &types.QuerySeckillActivityDetailResp{
		Code:    "000000",
		Message: "查询秒杀活动成功",
		Data:    data,
	}, nil
}
