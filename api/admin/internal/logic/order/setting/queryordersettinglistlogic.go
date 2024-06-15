package setting

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryOrderSettingListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryOrderSettingListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryOrderSettingListLogic {
	return &QueryOrderSettingListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryOrderSettingListLogic) QueryOrderSettingList(req *types.QueryOrderSettingListReq) (resp *types.QueryOrderSettingListResp, err error) {
	result, err := l.svcCtx.OrderSettingService.QueryOrderSettingList(l.ctx, &omsclient.QueryOrderSettingListReq{
		PageNum:  req.Current,
		PageSize: req.PageSize,
		Status:   req.Status,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询订单设置列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询订单设置失败")
	}

	var list []*types.QueryOrderSettingListData

	for _, item := range result.List {
		list = append(list, &types.QueryOrderSettingListData{
			CommentOvertime:     item.CommentOvertime,
			ConfirmOvertime:     item.ConfirmOvertime,
			FinishOvertime:      item.FinishOvertime,
			FlashOrderOvertime:  item.FlashOrderOvertime,
			Id:                  item.Id,
			IsDefault:           item.IsDefault,
			NormalOrderOvertime: item.NormalOrderOvertime,
			Status:              item.Status,
		})
	}

	return &types.QueryOrderSettingListResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    result.Total,
		Code:     "000000",
		Message:  "查询订单设置成功",
	}, nil
}
