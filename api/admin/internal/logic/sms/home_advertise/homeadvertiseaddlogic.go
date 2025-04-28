package home_advertise

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/common/res"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// HomeAdvertiseAddLogic 首页轮播广告
/*
Author: LiuFeiHua
Date: 2024/5/13 17:32
*/
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

// HomeAdvertiseAdd 添加首页轮播广告
func (l *HomeAdvertiseAddLogic) HomeAdvertiseAdd(req *types.AddHomeAdvertiseReq) (*types.BaseResp, error) {
	_, err := l.svcCtx.HomeAdvertiseService.AddHomeAdvertise(l.ctx, &smsclient.AddHomeAdvertiseReq{

		Name:      req.Name,                                                                          // 名称
		Type:      req.Type,                                                                          // 轮播位置：0->PC首页轮播；1->app首页轮播
		Pic:       "http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20181113/movie_ad.jpg", // 暂时没有上传,用这个当默认
		StartTime: req.StartTime,                                                                     // 开始时间
		EndTime:   req.EndTime,                                                                       // 结束时间
		Status:    req.Status,                                                                        // 上下线状态：0->下线；1->上线
		Url:       req.Url,                                                                           // 链接地址
		Note:      req.Note,                                                                          // 备注
		Sort:      req.Sort,                                                                          // 排序
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加首页广告信息失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return res.Success()
}
