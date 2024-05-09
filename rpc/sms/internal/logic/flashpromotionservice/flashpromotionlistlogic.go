package flashpromotionservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

type FlashPromotionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFlashPromotionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FlashPromotionListLogic {
	return &FlashPromotionListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FlashPromotionListLogic) FlashPromotionList(in *smsclient.FlashPromotionListReq) (*smsclient.FlashPromotionListResp, error) {
	q := query.SmsFlashPromotion.WithContext(l.ctx)
	if len(in.Title) > 0 {
		q = q.Where(query.SmsFlashPromotion.Title.Like("%" + in.Title + "%"))
	}

	if in.Status != 2 {
		q = q.Where(query.SmsFlashPromotion.Status.Eq(in.Status))
	}

	//if len(in.StartDate) > 0 {
	//		where = where + fmt.Sprintf(" AND start_date >= '%s'", in.StartDate)
	//	}
	//	if len(in.EndDate) > 0 {
	//		where = where + fmt.Sprintf(" AND end_date <= '%s'", in.EndDate)
	//	}
	offset := (in.Current - 1) * in.PageSize
	result, err := q.Offset(int(offset)).Limit(int(in.PageSize)).Find()
	count, err := q.Count()

	if err != nil {
		logc.Errorf(l.ctx, "查询限时购列表信息失败,参数：%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*smsclient.FlashPromotionListData
	for _, item := range result {

		list = append(list, &smsclient.FlashPromotionListData{
			Id:         item.ID,
			Title:      item.Title,
			StartDate:  item.StartDate.Format("2006-01-02"),
			EndDate:    item.EndDate.Format("2006-01-02"),
			Status:     item.Status,
			CreateTime: item.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	logc.Infof(l.ctx, "查询限时购列表信息,参数：%+v,响应：%+v", in, list)
	return &smsclient.FlashPromotionListResp{
		Total: count,
		List:  list,
	}, nil
}
