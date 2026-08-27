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

// UpdateHelpCategoryLogic 更新帮助分类
/*
Author: 刘飞华
Date: 2026/08/27 10:50:29
*/
type UpdateHelpCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateHelpCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateHelpCategoryLogic {
	return &UpdateHelpCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateHelpCategory 更新帮助分类
func (l *UpdateHelpCategoryLogic) UpdateHelpCategory(req *types.UpdateHelpCategoryReq) (resp *types.UpdateHelpCategoryResp, err error) {

	updateBy := l.ctx.Value("userName").(string)
	_, err = l.svcCtx.HelpCategoryService.UpdateHelpCategory(l.ctx, &cmsclient.UpdateHelpCategoryReq{
		Id:         req.Id,         // 主键id
		Name:       req.Name,       // 分类名称
		Icon:       req.Icon,       // 分类图标
		ShowStatus: req.ShowStatus, // 显示状态：0->不显示；1->显示
		Sort:       req.Sort,       // 排序
		UpdateBy:   updateBy,       // 更新者

	})

	if err != nil {
		logc.Errorf(l.ctx, "更新帮助分类失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.UpdateHelpCategoryResp{
		Code:    "000000",
		Message: "更新帮助分类成功",
	}, nil
}
