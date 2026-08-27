package help_category

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

// UpdateHelpCategoryStatusLogic 更新帮助分类状态状态
/*
Author: 刘飞华
Date: 2026/08/27 10:50:29
*/
type UpdateHelpCategoryStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateHelpCategoryStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateHelpCategoryStatusLogic {
	return &UpdateHelpCategoryStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateHelpCategoryStatus 更新帮助分类状态
func (l *UpdateHelpCategoryStatusLogic) UpdateHelpCategoryStatus(req *types.UpdateHelpCategoryStatusReq) (resp *types.UpdateHelpCategoryStatusResp, err error) {
	_, err = l.svcCtx.HelpCategoryService.UpdateHelpCategoryStatus(l.ctx, &cmsclient.UpdateHelpCategoryStatusReq{
		Ids:        req.Ids,        // 主键id
		ShowStatus: req.ShowStatus, // 显示状态：0->不显示；1->显示
		UpdateBy:   l.ctx.Value("userName").(string),
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新帮助分类状态失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.UpdateHelpCategoryStatusResp{
		Code:    "000000",
		Message: "更新帮助分类状态成功",
	}, nil
}
