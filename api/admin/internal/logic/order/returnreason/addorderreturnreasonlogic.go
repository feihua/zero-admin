package returnreason

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddOrderReturnReasonLogic 添加退货原因
/*
Author: LiuFeiHua
Date: 2024/6/15 11:40
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
func (l *AddOrderReturnReasonLogic) AddOrderReturnReason(req *types.AddOrderReturnReasonReq) (resp *types.AddOrderReturnReasonResp, err error) {
	_, err = l.svcCtx.OrderReturnReasonService.AddOrderReturnReason(l.ctx, &omsclient.AddOrderReturnReasonReq{
		Name:   req.Name,   // 退货类型
		Sort:   req.Sort,   // 排序
		Status: req.Status, // 状态：0->不启用；1->启用
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加退货原因地址信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("添加退货原因失败")
	}

	return &types.AddOrderReturnReasonResp{
		Code:    "000000",
		Message: "添加退货原因成功",
	}, nil
}
