package return_reason

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateOrderReturnReasonLogic 更新退货原因
/*
Author: LiuFeiHua
Date: 2025/05/26 15:21:44
*/
type UpdateOrderReturnReasonLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateOrderReturnReasonLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderReturnReasonLogic {
	return &UpdateOrderReturnReasonLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateOrderReturnReason 更新退货原因
func (l *UpdateOrderReturnReasonLogic) UpdateOrderReturnReason(req *types.UpdateOrderReturnReasonReq) (resp *types.BaseResp, err error) {
	userId, err := common.GetUserId(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.OrderReturnReasonService.UpdateOrderReturnReason(l.ctx, &omsclient.UpdateOrderReturnReasonReq{
		Id:       req.Id,     // 主键ID
		Name:     req.Name,   // 退货类型
		Sort:     req.Sort,   // 排序
		Status:   req.Status, // 状态：0->不启用；1->启用
		UpdateBy: userId,     // 更新人ID
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新退货原因失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "更新退货原因成功",
	}, nil
}
