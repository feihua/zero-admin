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

// AddOrderReturnReasonLogic 添加退货原因
/*
Author: LiuFeiHua
Date: 2025/05/26 15:21:44
*/
type AddOrderReturnReasonLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddOrderReturnReasonLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOrderReturnReasonLogic {
	return &AddOrderReturnReasonLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddOrderReturnReason 添加退货原因
func (l *AddOrderReturnReasonLogic) AddOrderReturnReason(req *types.AddOrderReturnReasonReq) (resp *types.BaseResp, err error) {
	userId, err := common.GetUserId(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.OrderReturnReasonService.AddOrderReturnReason(l.ctx, &omsclient.AddOrderReturnReasonReq{
		Name:     req.Name,   // 退货类型
		Sort:     req.Sort,   // 排序
		Status:   req.Status, // 状态：0->不启用；1->启用
		CreateBy: userId,     // 创建人ID
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加退货原因失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "添加退货原因成功",
	}, nil
}
