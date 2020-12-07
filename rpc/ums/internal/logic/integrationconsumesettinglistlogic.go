package logic

import (
	"context"
	"fmt"

	"go-zero-admin/rpc/ums/internal/svc"
	"go-zero-admin/rpc/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
)

type IntegrationConsumeSettingListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIntegrationConsumeSettingListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IntegrationConsumeSettingListLogic {
	return &IntegrationConsumeSettingListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IntegrationConsumeSettingListLogic) IntegrationConsumeSettingList(in *ums.IntegrationConsumeSettingListReq) (*ums.IntegrationConsumeSettingListResp, error) {
	all, _ := l.svcCtx.UmsIntegrationConsumeSettingModel.FindAll(in.Current, in.PageSize)
	//count, _ := l.svcCtx.UserModel.Count()

	var list []*ums.IntegrationConsumeSettingListData
	for _, item := range *all {

		list = append(list, &ums.IntegrationConsumeSettingListData{
			Id:                 item.Id,
			DeductionPerAmount: item.DeductionPerAmount,
			MaxPercentPerOrder: item.MaxPercentPerOrder,
			UseUnit:            item.UseUnit,
			CouponStatus:       item.CouponStatus,
		})
	}

	fmt.Println(list)
	return &ums.IntegrationConsumeSettingListResp{
		Total: 10,
		List:  list,
	}, nil

}
