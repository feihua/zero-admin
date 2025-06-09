package home_advertise

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

// QueryHomeAdvertiseDetailLogic 查询首页轮播广告详情
/*
Author: LiuFeiHua
Date: 2025/06/12 10:40:15
*/
type QueryHomeAdvertiseDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryHomeAdvertiseDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryHomeAdvertiseDetailLogic {
	return &QueryHomeAdvertiseDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryHomeAdvertiseDetail 查询首页轮播广告详情
func (l *QueryHomeAdvertiseDetailLogic) QueryHomeAdvertiseDetail(req *types.QueryHomeAdvertiseDetailReq) (resp *types.QueryHomeAdvertiseDetailResp, err error) {

	detail, err := l.svcCtx.HomeAdvertiseService.QueryHomeAdvertiseDetail(l.ctx, &smsclient.QueryHomeAdvertiseDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询首页轮播广告详情失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	data := types.QueryHomeAdvertiseDetailData{
		Id:         detail.Id,         // 编号
		Name:       detail.Name,       // 名称
		Type:       detail.Type,       // 轮播位置：0->PC首页轮播；1->app首页轮播
		Pic:        detail.Pic,        // 图片地址
		StartTime:  detail.StartTime,  // 开始时间
		EndTime:    detail.EndTime,    // 结束时间
		Status:     detail.Status,     // 上下线状态：0->下线；1->上线
		ClickCount: detail.ClickCount, // 点击数
		OrderCount: detail.OrderCount, // 下单数
		Url:        detail.Url,        // 链接地址
		Remark:     detail.Remark,     // 备注
		Sort:       detail.Sort,       // 排序
		CreateTime: detail.CreateTime, // 创建时间
		UpdateTime: detail.UpdateTime, // 更新时间

	}
	return &types.QueryHomeAdvertiseDetailResp{
		Code:    "000000",
		Message: "查询首页轮播广告成功",
		Data:    data,
	}, nil
}
