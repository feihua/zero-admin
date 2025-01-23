package homeadvertiseservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// QueryHomeAdvertiseDetailLogic 查询首页轮播广告详情
/*
Author: LiuFeiHua
Date: 2025/01/23 16:18:56
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

	if err != nil {
		logc.Errorf(l.ctx, "查询首页轮播广告详情失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询首页轮播广告详情失败")
	}

	data := &smsclient.QueryHomeAdvertiseDetailResp{
		Id:         item.ID,                                      //
		Name:       item.Name,                                    // 名称
		Type:       item.Type,                                    // 轮播位置：0->PC首页轮播；1->app首页轮播
		Pic:        item.Pic,                                     // 图片地址
		StartTime:  item.StartTime.Format("2006-01-02 15:04:05"), // 开始时间
		EndTime:    item.EndTime.Format("2006-01-02 15:04:05"),   // 结束时间
		Status:     item.Status,                                  // 上下线状态：0->下线；1->上线
		ClickCount: item.ClickCount,                              // 点击数
		OrderCount: item.OrderCount,                              // 下单数
		Url:        item.URL,                                     // 链接地址
		Note:       item.Note,                                    // 备注
		Sort:       item.Sort,                                    // 排序
	}

	logc.Infof(l.ctx, "查询首页轮播广告详情,参数：%+v,响应：%+v", in, data)
	return data, nil
}
