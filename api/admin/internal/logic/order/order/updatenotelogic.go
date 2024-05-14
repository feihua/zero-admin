package order

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateNoteLogic 订单信息
/*
Author: LiuFeiHua
Date: 2024/5/14 17:05
*/
type UpdateNoteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateNoteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateNoteLogic {
	return &UpdateNoteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateNote 备注订单
func (l *UpdateNoteLogic) UpdateNote(req *types.UpdateNoteReq) (resp *types.UpdateNoteResp, err error) {
	_, err = l.svcCtx.OrderService.UpdateNote(l.ctx, &omsclient.UpdateNoteReq{
		OrderId:    req.Id,
		Status:     req.Status,
		Note:       req.Note,
		OperateMan: l.ctx.Value("userName").(string),
	})

	if err != nil {
		logc.Errorf(l.ctx, "备注订单失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("备注订单失败")
	}

	return &types.UpdateNoteResp{
		Code:    "000000",
		Message: "备注订单成功",
	}, nil
}
