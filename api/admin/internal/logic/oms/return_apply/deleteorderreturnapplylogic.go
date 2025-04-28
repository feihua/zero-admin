package return_apply

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/common/res"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteOrderReturnApplyLogic 删除订单退货申请
/*
Author: LiuFeiHua
Date: 2024/6/15 11:36
*/
type DeleteOrderReturnApplyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteOrderReturnApplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteOrderReturnApplyLogic {
	return &DeleteOrderReturnApplyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteOrderReturnApply 删除订单退货申请
func (l *DeleteOrderReturnApplyLogic) DeleteOrderReturnApply(req *types.DeleteOrderReturnApplyReq) (resp *types.BaseResp, err error) {
	_, err = l.svcCtx.OrderReturnApplyService.DeleteOrderReturnApply(l.ctx, &omsclient.DeleteOrderReturnApplyReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据Id: %+v,删除退货申请异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("删除退货申请失败")
	}

	return res.Success()
}
