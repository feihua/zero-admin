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

// UpdateHomeAdvertiseLogic 更新首页轮播广告表
/*
Author: LiuFeiHua
Date: 2024/6/12 17:51
*/
type UpdateHomeAdvertiseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateHomeAdvertiseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateHomeAdvertiseLogic {
	return &UpdateHomeAdvertiseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateHomeAdvertise 更新首页轮播广告表
func (l *UpdateHomeAdvertiseLogic) UpdateHomeAdvertise(in *smsclient.UpdateHomeAdvertiseReq) (*smsclient.UpdateHomeAdvertiseResp, error) {
	StartTime, _ := time.Parse("2006-01-02 15:04:05", in.StartTime)
	EndTime, _ := time.Parse("2006-01-02 15:04:05", in.EndTime)

	q := query.SmsHomeAdvertise
	_, err := q.WithContext(l.ctx).Updates(&model.SmsHomeAdvertise{
		ID:         in.Id,
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

	return &smsclient.UpdateHomeAdvertiseResp{}, nil
}
