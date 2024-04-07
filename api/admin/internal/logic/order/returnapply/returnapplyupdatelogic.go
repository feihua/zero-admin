package logic

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReturnApplyUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReturnApplyUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) ReturnApplyUpdateLogic {
	return ReturnApplyUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReturnApplyUpdateLogic) ReturnApplyUpdate(req types.UpdateReturnApplyReq) (*types.UpdateReturnApplyResp, error) {
	userName := l.ctx.Value("userName").(string)
	_, err := l.svcCtx.OrderReturnApplyService.OrderReturnApplyUpdate(l.ctx, &omsclient.OrderReturnApplyUpdateReq{
		Id:               req.Id,
		CompanyAddressId: req.CompanyAddressId,
		Status:           req.Status,
		HandleNote:       req.HandleNote,
		HandleMan:        userName,
		ReceiveMan:       userName,
		ReceiveNote:      req.ReceiveNote,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("更新退货申请信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("更新退货申请失败")
	}

	return &types.UpdateReturnApplyResp{
		Code:    "000000",
		Message: "更新退货申请成功",
	}, nil
}
