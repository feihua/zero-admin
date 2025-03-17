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

// QueryOrderReturnReasonDetailLogic 查询退货原因表详情
/*
Author: 刘飞华
Date: 2025/02/05 11:40:41
*/
type QueryOrderReturnReasonDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryOrderReturnReasonDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryOrderReturnReasonDetailLogic {
	return &QueryOrderReturnReasonDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryOrderReturnReasonDetail 查询退货原因表详情
func (l *QueryOrderReturnReasonDetailLogic) QueryOrderReturnReasonDetail(req *types.QueryOrderReturnReasonDetailReq) (resp *types.QueryOrderReturnReasonDetailResp, err error) {

	detail, err := l.svcCtx.OrderReturnReasonService.QueryOrderReturnReasonDetail(l.ctx, &omsclient.QueryOrderReturnReasonDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询退货原因表详情失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	data := types.QueryOrderReturnReasonDetailData{
		Id:         detail.Id,         //
		Name:       detail.Name,       // 退货类型
		Sort:       detail.Sort,       // 排序
		Status:     detail.Status,     // 状态：0->不启用；1->启用
		CreateTime: detail.CreateTime, // 创建时间
	}
	return &types.QueryOrderReturnReasonDetailResp{
		Code:    "000000",
		Message: "查询退货原因表成功",
		Data:    data,
	}, nil
}
