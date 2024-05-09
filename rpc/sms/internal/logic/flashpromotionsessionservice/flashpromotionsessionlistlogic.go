package flashpromotionsessionservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FlashPromotionSessionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFlashPromotionSessionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlashPromotionSessionListLogic {
	return &FlashPromotionSessionListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FlashPromotionSessionListLogic) FlashPromotionSessionList(in *smsclient.FlashPromotionSessionListReq) (*smsclient.FlashPromotionSessionListResp, error) {
	q := query.SmsFlashPromotionSession.WithContext(l.ctx)

	offset := (in.Current - 1) * in.PageSize
	result, err := q.Offset(int(offset)).Limit(int(in.PageSize)).Find()
	count, err := q.Count()

	if err != nil {
		logc.Errorf(l.ctx, "查询限时购场次列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*smsclient.FlashPromotionSessionListData
	for _, item := range result {

		list = append(list, &smsclient.FlashPromotionSessionListData{
			Id:         item.ID,
			Name:       item.Name,
			StartTime:  item.StartTime.Format("2006-01-02 15:04:05"),
			EndTime:    item.EndTime.Format("2006-01-02 15:04:05"),
			Status:     item.Status,
			CreateTime: item.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	logc.Infof(l.ctx, "查询限时购场次列表信息,参数：%+v,响应：%+v", in, list)
	return &smsclient.FlashPromotionSessionListResp{
		Total: count,
		List:  list,
	}, nil
}
