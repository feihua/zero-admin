package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReturnResonListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReturnResonListLogic(ctx context.Context, svcCtx *svc.ServiceContext) ReturnResonListLogic {
	return ReturnResonListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReturnResonListLogic) ReturnResonList(req types.ListReturnResonReq) (*types.ListReturnResonResp, error) {
	resp, err := l.svcCtx.OrderReturnReasonService.OrderReturnReasonList(l.ctx, &omsclient.OrderReturnReasonListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
		Name:     req.Name,
		Status:   req.Status,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询退货原因列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询退货原因失败")
	}

	var list []*types.ListReturnResonData

	for _, item := range resp.List {
		list = append(list, &types.ListReturnResonData{
			Id:         item.Id,
			Name:       item.Name,
			Sort:       item.Sort,
			Status:     item.Status,
			CreateTime: item.CreateTime,
		})
	}

	return &types.ListReturnResonResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
		Code:     "000000",
		Message:  "查询退货原因成功",
	}, nil
}
