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

// QueryOrderReturnReasonDetailLogic 查询退货原因详情
/*
Author: LiuFeiHua
Date: 2025/05/26 15:21:44
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

// QueryOrderReturnReasonDetail 查询退货原因详情
func (l *QueryOrderReturnReasonDetailLogic) QueryOrderReturnReasonDetail(req *types.QueryOrderReturnReasonDetailReq) (resp *types.QueryOrderReturnReasonDetailResp, err error) {

	detail, err := l.svcCtx.OrderReturnReasonService.QueryOrderReturnReasonDetail(l.ctx, &omsclient.QueryOrderReturnReasonDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询退货原因详情失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	data := types.QueryOrderReturnReasonDetailData{
		Id:         detail.Id,         // 主键ID
		Name:       detail.Name,       // 退货类型
		Sort:       detail.Sort,       // 排序
		Status:     detail.Status,     // 状态：0->不启用；1->启用
		CreateBy:   detail.CreateBy,   // 创建人ID
		CreateTime: detail.CreateTime, // 创建时间
		UpdateBy:   detail.UpdateBy,   // 更新人ID
		UpdateTime: detail.UpdateTime, // 更新时间

	}
	return &types.QueryOrderReturnReasonDetailResp{
		Code:    "000000",
		Message: "查询退货原因成功",
		Data:    data,
	}, nil
}
