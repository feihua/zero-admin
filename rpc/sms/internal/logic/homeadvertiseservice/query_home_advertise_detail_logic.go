package homeadvertiseservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

// QueryHomeAdvertiseDetailLogic 查询首页轮播广告详情
/*
Author: LiuFeiHua
Date: 2025/06/11 10:44:50
*/
type QueryHomeAdvertiseDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryHomeAdvertiseDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryHomeAdvertiseDetailLogic {
	return &QueryHomeAdvertiseDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryHomeAdvertiseDetail 查询首页轮播广告详情
func (l *QueryHomeAdvertiseDetailLogic) QueryHomeAdvertiseDetail(in *smsclient.QueryHomeAdvertiseDetailReq) (*smsclient.QueryHomeAdvertiseDetailResp, error) {
	item, err := query.SmsHomeAdvertise.WithContext(l.ctx).Where(query.SmsHomeAdvertise.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "首页轮播广告不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("首页轮播广告不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询首页轮播广告异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询首页轮播广告异常")
	}

	data := &smsclient.QueryHomeAdvertiseDetailResp{
		Id:         item.ID,                                 // 编号
		Name:       item.Name,                               // 名称
		Type:       item.Type,                               // 轮播位置：0->PC首页轮播；1->app首页轮播
		Pic:        item.Pic,                                // 图片地址
		StartTime:  time_util.TimeToStr(item.StartTime),     // 开始时间
		EndTime:    time_util.TimeToStr(item.EndTime),       // 结束时间
		Status:     item.Status,                             // 上下线状态：0->下线；1->上线
		ClickCount: item.ClickCount,                         // 点击数
		OrderCount: item.OrderCount,                         // 下单数
		Url:        item.URL,                                // 链接地址
		Remark:     item.Remark,                             // 备注
		Sort:       item.Sort,                               // 排序
		CreateTime: time_util.TimeToStr(item.CreateTime),    // 创建时间
		UpdateTime: time_util.TimeToString(item.UpdateTime), // 更新时间
	}

	return data, nil
}
