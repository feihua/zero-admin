package return_apply

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateOrderReturnApplyStatusLogic 修改订单退货申请状态
/*
Author: LiuFeiHua
Date: 2024/6/15 11:38
*/
type UpdateOrderReturnApplyStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateOrderReturnApplyStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderReturnApplyStatusLogic {
	return &UpdateOrderReturnApplyStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateOrderReturnApplyStatus 修改订单退货申请状态
/*
Author: LiuFeiHua
Date: 2024/6/15 11:38
*/
func (l *UpdateOrderReturnApplyStatusLogic) UpdateOrderReturnApplyStatus(req *types.UpdateOrderReturnApplyStatusReq) (resp *types.UpdateOrderReturnApplyStatusResp, err error) {
	userName := l.ctx.Value("userName").(string)
	_, err = l.svcCtx.OrderReturnApplyService.UpdateOrderReturnApply(l.ctx, &omsclient.UpdateOrderReturnApplyReq{
		Id:               req.Id,
		CompanyAddressId: req.CompanyAddressId,
		Status:           req.Status,
		HandleNote:       req.HandleNote,
		HandleMan:        userName,
		ReceiveMan:       userName,
		ReceiveNote:      req.ReceiveNote,
		ReturnAmount:     req.ReturnAmount,
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新退货申请信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("更新退货申请失败")
	}

	return &types.UpdateOrderReturnApplyStatusResp{
		Code:    "000000",
		Message: "更新退货申请成功",
	}, nil
}
