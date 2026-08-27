package help

import (
	"context"

	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateHelpStatusLogic 更新帮助状态状态
/*
Author: 刘飞华
Date: 2026/08/27 10:50:29
*/
type UpdateHelpStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateHelpStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateHelpStatusLogic {
	return &UpdateHelpStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateHelpStatus 更新帮助状态
func (l *UpdateHelpStatusLogic) UpdateHelpStatus(req *types.UpdateHelpStatusReq) (resp *types.UpdateHelpStatusResp, err error) {
	_, err = l.svcCtx.HelpService.UpdateHelpStatus(l.ctx, &cmsclient.UpdateHelpStatusReq{
		Ids:        req.Ids,        // 主键id
		ShowStatus: req.ShowStatus, // 显示状态：0->不显示；1->显示
		UpdateBy:   l.ctx.Value("userName").(string),
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新帮助状态失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.UpdateHelpStatusResp{
		Code:    "000000",
		Message: "更新帮助状态成功",
	}, nil
}
