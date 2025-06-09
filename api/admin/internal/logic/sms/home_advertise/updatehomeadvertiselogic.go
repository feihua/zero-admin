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

// UpdateHomeAdvertiseLogic 更新首页轮播广告
/*
Author: LiuFeiHua
Date: 2025/06/12 10:40:15
*/
type UpdateHomeAdvertiseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateHomeAdvertiseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateHomeAdvertiseLogic {
	return &UpdateHomeAdvertiseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateHomeAdvertise 更新首页轮播广告
func (l *UpdateHomeAdvertiseLogic) UpdateHomeAdvertise(req *types.UpdateHomeAdvertiseReq) (resp *types.BaseResp, err error) {

	_, err = l.svcCtx.HomeAdvertiseService.UpdateHomeAdvertise(l.ctx, &smsclient.UpdateHomeAdvertiseReq{
		Id:        req.Id,        // 编号
		Name:      req.Name,      // 名称
		Type:      req.Type,      // 轮播位置：0->PC首页轮播；1->app首页轮播
		Pic:       req.Pic,       // 图片地址
		StartTime: req.StartTime, // 开始时间
		EndTime:   req.EndTime,   // 结束时间
		Status:    req.Status,    // 上下线状态：0->下线；1->上线
		Url:       req.Url,       // 链接地址
		Remark:    req.Remark,    // 备注
		Sort:      req.Sort,      // 排序
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新首页轮播广告失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "更新首页轮播广告成功",
	}, nil
}
