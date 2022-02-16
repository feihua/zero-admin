package logic

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/sms/internal/svc"
	"zero-admin/rpc/sms/sms"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *HomeAdvertiseListLogic) HomeAdvertiseList(in *sms.HomeAdvertiseListReq) (*sms.HomeAdvertiseListResp, error) {
	all, err := l.svcCtx.SmsHomeAdvertiseModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.SmsHomeAdvertiseModel.Count()

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询首页广告列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	var list []*sms.HomeAdvertiseListData
	for _, item := range *all {

		list = append(list, &sms.HomeAdvertiseListData{
			Id:         item.Id,
			Name:       item.Name,
			Type:       item.Type,
			Pic:        item.Pic,
			StartTime:  item.StartTime.Format("2006-01-02 15:04:05"),
			EndTime:    item.EndTime.Format("2006-01-02 15:04:05"),
			Status:     item.Status,
			ClickCount: item.ClickCount,
			OrderCount: item.OrderCount,
			Url:        item.Url,
			Note:       item.Note,
			Sort:       item.Sort,
		})
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logx.WithContext(l.ctx).Infof("查询首页广告列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &sms.HomeAdvertiseListResp{
		Total: count,
		List:  list,
	}, nil
}
