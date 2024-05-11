package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *HomeAdvertiseAddLogic) HomeAdvertiseAdd(req types.AddHomeAdvertiseReq) (*types.AddHomeAdvertiseResp, error) {
	_, err := l.svcCtx.HomeAdvertiseService.HomeAdvertiseAdd(l.ctx, &smsclient.HomeAdvertiseAddReq{
		Name:      req.Name,
		Type:      req.Type,
		Pic:       "http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20181113/movie_ad.jpg", //暂时没有上传,用这个当默认
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		Status:    req.Status,
		Url:       req.Url,
		Note:      req.Note,
		Sort:      req.Sort,
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加首页广告信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("添加首页广告失败")
	}

	return &types.AddHomeAdvertiseResp{
		Code:    "000000",
		Message: "添加首页广告成功",
	}, nil
}
