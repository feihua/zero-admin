package homeadvertiseservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"time"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddHomeAdvertiseLogic 添加首页轮播广告表
/*
Author: LiuFeiHua
Date: 2024/6/12 17:50
*/
type AddHomeAdvertiseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddHomeAdvertiseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddHomeAdvertiseLogic {
	return &AddHomeAdvertiseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddHomeAdvertise 添加首页轮播广告表
func (l *AddHomeAdvertiseLogic) AddHomeAdvertise(in *smsclient.AddHomeAdvertiseReq) (*smsclient.AddHomeAdvertiseResp, error) {
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

	return &smsclient.AddHomeAdvertiseResp{}, nil
}
