package integrationconsumesetting

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryIntegrationConsumeSettingListLogic 查询积分消费设置列表
/*
Author: LiuFeiHua
Date: 2024/5/23 9:38
*/
type QueryIntegrationConsumeSettingListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryIntegrationConsumeSettingListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryIntegrationConsumeSettingListLogic {
	return &QueryIntegrationConsumeSettingListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryIntegrationConsumeSettingList 查询积分消费设置列表
func (l *QueryIntegrationConsumeSettingListLogic) QueryIntegrationConsumeSettingList(req *types.QueryIntegrationConsumeSettingListReq) (resp *types.QueryIntegrationConsumeSettingListResp, err error) {
	result, err := l.svcCtx.IntegrationConsumeSettingService.QueryIntegrationConsumeSettingList(l.ctx, &umsclient.QueryIntegrationConsumeSettingListReq{
		PageNum:  req.Current,
		PageSize: req.PageSize,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询积分消费设置列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询积分消费设置失败")
	}

	var list []*types.QueryIntegrationConsumeSettingListData

	for _, item := range result.List {
		list = append(list, &types.QueryIntegrationConsumeSettingListData{
			CouponStatus:       item.CouponStatus,
			DeductionPerAmount: item.DeductionPerAmount,
			Id:                 item.Id,
			IsDefault:          item.IsDefault,
			MaxPercentPerOrder: item.MaxPercentPerOrder,
			UseUnit:            item.UseUnit,
		})
	}

	return &types.QueryIntegrationConsumeSettingListResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    result.Total,
		Code:     "000000",
		Message:  "查询积分消费设置成功",
	}, nil
}
