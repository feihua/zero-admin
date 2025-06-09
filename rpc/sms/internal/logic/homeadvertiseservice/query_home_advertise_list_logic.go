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
	"time"
)

// QueryHomeAdvertiseListLogic 查询首页轮播广告列表
/*
Author: LiuFeiHua
Date: 2025/06/11 10:44:50
*/
type QueryHomeAdvertiseListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryHomeAdvertiseListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryHomeAdvertiseListLogic {
	return &QueryHomeAdvertiseListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryHomeAdvertiseList 查询首页轮播广告列表
func (l *QueryHomeAdvertiseListLogic) QueryHomeAdvertiseList(in *smsclient.QueryHomeAdvertiseListReq) (*smsclient.QueryHomeAdvertiseListResp, error) {
	homeAdvertise := query.SmsHomeAdvertise
	q := homeAdvertise.WithContext(l.ctx)
	if len(in.Name) > 0 {
		q = q.Where(homeAdvertise.Name.Like("%" + in.Name + "%"))
	}
	if in.Type != 2 {
		q = q.Where(homeAdvertise.Type.Eq(in.Type))
	}

	if len(in.StartTime) > 0 {
		startTime, _ := time.Parse("2006-01-02 15:04:05", in.StartTime)
		q = q.Where(query.SmsHomeAdvertise.StartTime.Gte(startTime))
	}
	if len(in.EndTime) > 0 {
		endTime, _ := time.Parse("2006-01-02 15:04:05", in.EndTime)
		q = q.Where(query.SmsHomeAdvertise.EndTime.Lte(endTime))
	}
	if in.Status != 2 {
		q = q.Where(homeAdvertise.Status.Eq(in.Status))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询首页轮播广告列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询首页轮播广告列表失败")
	}

	var list []*smsclient.HomeAdvertiseListData

	for _, item := range result {
		list = append(list, &smsclient.HomeAdvertiseListData{
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

		})
	}

	return &smsclient.QueryHomeAdvertiseListResp{
		Total: count,
		List:  list,
	}, nil
}
