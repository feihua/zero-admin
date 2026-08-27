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

// AddHelpLogic 添加帮助
/*
Author: 刘飞华
Date: 2026/08/27 10:50:29
*/
type AddHelpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddHelpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddHelpLogic {
	return &AddHelpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddHelp 添加帮助
func (l *AddHelpLogic) AddHelp(req *types.AddHelpReq) (resp *types.AddHelpResp, err error) {
	createBy := l.ctx.Value("userName").(string)

	_, err = l.svcCtx.HelpService.AddHelp(l.ctx, &cmsclient.AddHelpReq{
		CategoryId: req.CategoryId, // 分类id
		Icon:       req.Icon,       // 图标
		Title:      req.Title,      // 标题
		ShowStatus: req.ShowStatus, // 显示状态：0->不显示；1->显示
		ReadCount:  0,              // 阅读量
		Content:    req.Content,    // 内容
		CreateBy:   createBy,       // 创建者

	})

	if err != nil {
		logc.Errorf(l.ctx, "添加帮助失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.AddHelpResp{
		Code:    "000000",
		Message: "添加帮助成功",
	}, nil
}
