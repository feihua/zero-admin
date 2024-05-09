package homeadvertiseservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"

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

	//if len(in.StartTime) > 0 {
	//		where = where + fmt.Sprintf(" AND start_time >= '%s'", in.StartTime)
	//	}
	//	if len(in.EndTime) > 0 {
	//		where = where + fmt.Sprintf(" AND end_time <= '%s'", in.EndTime)
	//	}
	offset := (in.Current - 1) * in.PageSize
	result, err := q.Offset(int(offset)).Limit(int(in.PageSize)).Find()
	count, err := q.Count()

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
