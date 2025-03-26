package homeadvertiseservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"
	"time"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddHomeAdvertiseLogic 添加首页轮播广告
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

// AddHomeAdvertise 添加首页轮播广告
func (l *AddHomeAdvertiseLogic) AddHomeAdvertise(in *smsclient.AddHomeAdvertiseReq) (*smsclient.AddHomeAdvertiseResp, error) {
	StartTime, _ := time.Parse("2006-01-02 15:04:05", in.StartTime)
	EndTime, _ := time.Parse("2006-01-02 15:04:05", in.EndTime)

	err := query.SmsHomeAdvertise.WithContext(l.ctx).Create(&model.SmsHomeAdvertise{
		Name:       in.Name,       // 名称
		Type:       in.Type,       // 轮播位置：0->PC首页轮播；1->app首页轮播
		Pic:        in.Pic,        // 图片地址
		StartTime:  StartTime,     // 开始时间
		EndTime:    EndTime,       // 结束时间
		Status:     in.Status,     // 上下线状态：0->下线；1->上线
		ClickCount: in.ClickCount, // 点击数
		OrderCount: in.OrderCount, // 下单数
		URL:        in.Url,        // 链接地址
		Note:       in.Note,       // 备注
		Sort:       in.Sort,       // 排序
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加首页轮播广告失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("添加首页轮播广告失败")
	}

	return &smsclient.AddHomeAdvertiseResp{}, nil
}
