package homeadvertiseservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"time"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// HomeAdvertiseAddLogic 首页广告
/*
Author: LiuFeiHua
Date: 2024/5/13 17:36
*/
type HomeAdvertiseAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomeAdvertiseAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeAdvertiseAddLogic {
	return &HomeAdvertiseAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// HomeAdvertiseAdd 添加首页广告
func (l *HomeAdvertiseAddLogic) HomeAdvertiseAdd(in *smsclient.HomeAdvertiseAddReq) (*smsclient.HomeAdvertiseAddResp, error) {
	StartTime, _ := time.Parse("2006-01-02 15:04:05", in.StartTime)
	EndTime, _ := time.Parse("2006-01-02 15:04:05", in.EndTime)

	err := query.SmsHomeAdvertise.WithContext(l.ctx).Create(&model.SmsHomeAdvertise{
		Name:       in.Name,
		Type:       in.Type,
		Pic:        in.Pic,
		StartTime:  StartTime,
		EndTime:    EndTime,
		Status:     in.Status,
		ClickCount: in.ClickCount,
		OrderCount: in.OrderCount,
		URL:        in.Url,
		Note:       &in.Note,
		Sort:       in.Sort,
	})

	if err != nil {
		return nil, err
	}
	return &smsclient.HomeAdvertiseAddResp{}, nil
}
