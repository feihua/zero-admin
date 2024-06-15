package integrationconsumesettingservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryIntegrationConsumeSettingListLogic 查询积分消费设置列表
/*
Author: LiuFeiHua
Date: 2024/6/11 14:21
*/
type QueryIntegrationConsumeSettingListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryIntegrationConsumeSettingListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryIntegrationConsumeSettingListLogic {
	return &QueryIntegrationConsumeSettingListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryIntegrationConsumeSettingList 查询积分消费设置列表
func (l *QueryIntegrationConsumeSettingListLogic) QueryIntegrationConsumeSettingList(in *umsclient.QueryIntegrationConsumeSettingListReq) (*umsclient.QueryIntegrationConsumeSettingListResp, error) {
	q := query.UmsIntegrationConsumeSetting.WithContext(l.ctx)
	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询积分消费设置列表信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*umsclient.IntegrationConsumeSettingListData
	for _, item := range result {

		list = append(list, &umsclient.IntegrationConsumeSettingListData{
			Id:                 item.ID,
			DeductionPerAmount: item.DeductionPerAmount,
			MaxPercentPerOrder: item.MaxPercentPerOrder,
			UseUnit:            item.UseUnit,
			CouponStatus:       item.CouponStatus,
			IsDefault:          item.IsDefault,
		})
	}

	return &umsclient.QueryIntegrationConsumeSettingListResp{
		Total: count,
		List:  list,
	}, nil

}
