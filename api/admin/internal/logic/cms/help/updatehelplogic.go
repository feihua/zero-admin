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

// UpdateHelpLogic 更新帮助
/*
Author: 刘飞华
Date: 2026/08/27 10:50:29
*/
type UpdateHelpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateHelpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateHelpLogic {
	return &UpdateHelpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateHelp 更新帮助
func (l *UpdateHelpLogic) UpdateHelp(req *types.UpdateHelpReq) (resp *types.UpdateHelpResp, err error) {

	updateBy := l.ctx.Value("userName").(string)
	_, err = l.svcCtx.HelpService.UpdateHelp(l.ctx, &cmsclient.UpdateHelpReq{
		Id:         req.Id,         // 主键id
		CategoryId: req.CategoryId, // 分类id
		Icon:       req.Icon,       // 图标
		Title:      req.Title,      // 标题
		ShowStatus: req.ShowStatus, // 显示状态：0->不显示；1->显示
		Content:    req.Content,    // 内容
		UpdateBy:   updateBy,       // 更新者

	})

	if err != nil {
		logc.Errorf(l.ctx, "更新帮助失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.UpdateHelpResp{
		Code:    "000000",
		Message: "更新帮助成功",
	}, nil
}
