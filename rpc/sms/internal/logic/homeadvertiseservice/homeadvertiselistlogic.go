package homeadvertiseservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

// HomeAdvertiseListLogic 首页广告
/*
Author: LiuFeiHua
Date: 2024/5/13 17:37
*/
type HomeAdvertiseListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomeAdvertiseListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeAdvertiseListLogic {
	return &HomeAdvertiseListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// HomeAdvertiseList 查询首页广告
func (l *HomeAdvertiseListLogic) HomeAdvertiseList(in *smsclient.HomeAdvertiseListReq) (*smsclient.HomeAdvertiseListResp, error) {
	q := query.SmsHomeAdvertise.WithContext(l.ctx)
	if len(in.Name) > 0 {
		q = q.Where(query.SmsHomeAdvertise.Name.Like("%" + in.Name + "%"))
	}

	if in.Type != 2 {
		q = q.Where(query.SmsHomeAdvertise.Type.Eq(in.Type))
	}
	if in.Status != 2 {
		q = q.Where(query.SmsHomeAdvertise.Status.Eq(in.Status))
	}

	if len(in.StartTime) > 0 {
		startTime, _ := time.Parse("2006-01-02 15:04:05", in.StartTime)
		q = q.Where(query.SmsHomeAdvertise.StartTime.Gte(startTime))
	}
	if len(in.EndTime) > 0 {
		endTime, _ := time.Parse("2006-01-02 15:04:05", in.EndTime)
		q = q.Where(query.SmsHomeAdvertise.EndTime.Lte(endTime))
	}
	result, count, err := q.FindByPage(int((in.Current-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询首页广告列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*smsclient.HomeAdvertiseListData
	for _, item := range result {

		list = append(list, &smsclient.HomeAdvertiseListData{
			Id:         item.ID,
			Name:       item.Name,
			Type:       item.Type,
			Pic:        item.Pic,
			StartTime:  item.StartTime.Format("2006-01-02 15:04:05"),
			EndTime:    item.EndTime.Format("2006-01-02 15:04:05"),
			Status:     item.Status,
			ClickCount: item.ClickCount,
			OrderCount: item.OrderCount,
			Url:        item.URL,
			Note:       *item.Note,
			Sort:       item.Sort,
		})
	}

	logc.Infof(l.ctx, "查询首页广告列表信息,参数：%+v,响应：%+v", in, list)
	return &smsclient.HomeAdvertiseListResp{
		Total: count,
		List:  list,
	}, nil
}
