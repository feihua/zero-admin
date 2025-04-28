package order

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

// CloseOrderLogic 订单信息
/*
Author: LiuFeiHua
Date: 2024/5/14 17:04
*/
type CloseOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCloseOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CloseOrderLogic {
	return &CloseOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CloseOrder 批量关闭订单
func (l *CloseOrderLogic) CloseOrder(req *types.CloseOrderReq) (resp *types.BaseResp, err error) {
	_, err = l.svcCtx.OrderService.CloseOrder(l.ctx, &omsclient.CloseOrderReq{
		Ids:        req.Ids,
		Note:       req.Note,
		OperateMan: l.ctx.Value("userName").(string),
	})

	if err != nil {
		logc.Errorf(l.ctx, "批量关闭订单失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("批量关闭订单失败")
	}

	return res.Success()
}
