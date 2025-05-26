package return_reason

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryOrderReturnReasonListLogic 查询退货原因列表
/*
Author: LiuFeiHua
Date: 2025/05/26 15:21:44
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

// QueryOrderReturnReasonList 查询退货原因列表
func (l *QueryOrderReturnReasonListLogic) QueryOrderReturnReasonList(req *types.QueryOrderReturnReasonListReq) (resp *types.QueryOrderReturnReasonListResp, err error) {
	result, err := l.svcCtx.OrderReturnReasonService.QueryOrderReturnReasonList(l.ctx, &omsclient.QueryOrderReturnReasonListReq{
		PageNum:  req.Current,
		PageSize: req.PageSize,
		Name:     req.Name,   // 退货类型
		Status:   req.Status, // 状态：0->不启用；1->启用
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询字退货原因列表失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.QueryOrderReturnReasonListData

	for _, detail := range result.List {
		list = append(list, &types.QueryOrderReturnReasonListData{
			Id:         detail.Id,         // 主键ID
			Name:       detail.Name,       // 退货类型
			Sort:       detail.Sort,       // 排序
			Status:     detail.Status,     // 状态：0->不启用；1->启用
			CreateBy:   detail.CreateBy,   // 创建人ID
			CreateTime: detail.CreateTime, // 创建时间
			UpdateBy:   detail.UpdateBy,   // 更新人ID
			UpdateTime: detail.UpdateTime, // 更新时间

		})
	}

	return &types.QueryOrderReturnReasonListResp{
		Code:     "000000",
		Message:  "查询退货原因列表成功",
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    result.Total,
	}, nil
}
