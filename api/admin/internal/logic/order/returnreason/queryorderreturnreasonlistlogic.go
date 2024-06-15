package returnreason

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"strings"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryOrderReturnReasonListLogic 查询退货原因
/*
Author: LiuFeiHua
Date: 2024/6/15 11:41
*/
type QueryOrderReturnReasonListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryOrderReturnReasonListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryOrderReturnReasonListLogic {
	return &QueryOrderReturnReasonListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryOrderReturnReasonList 查询退货原因
func (l *QueryOrderReturnReasonListLogic) QueryOrderReturnReasonList(req *types.QueryOrderReturnReasonListReq) (resp *types.QueryOrderReturnReasonListResp, err error) {
	result, err := l.svcCtx.OrderReturnReasonService.QueryOrderReturnReasonList(l.ctx, &omsclient.QueryOrderReturnReasonListReq{
		PageNum:  req.Current,
		PageSize: req.PageSize,
		Name:     strings.TrimSpace(req.Name),
		Status:   req.Status,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询退货原因列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询退货原因失败")
	}

	var list []*types.QueryOrderReturnReasonListData

	for _, item := range result.List {
		list = append(list, &types.QueryOrderReturnReasonListData{
			Id:         item.Id,
			Name:       item.Name,
			Sort:       item.Sort,
			Status:     item.Status,
			CreateTime: item.CreateTime,
		})
	}

	return &types.QueryOrderReturnReasonListResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    result.Total,
		Code:     "000000",
		Message:  "查询退货原因成功",
	}, nil
}
