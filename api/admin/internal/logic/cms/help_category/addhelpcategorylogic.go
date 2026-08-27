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

// AddHelpCategoryLogic 添加帮助分类
/*
Author: 刘飞华
Date: 2026/08/27 10:50:29
*/
type AddHelpCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddHelpCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddHelpCategoryLogic {
	return &AddHelpCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddHelpCategory 添加帮助分类
func (l *AddHelpCategoryLogic) AddHelpCategory(req *types.AddHelpCategoryReq) (resp *types.AddHelpCategoryResp, err error) {
	createBy := l.ctx.Value("userName").(string)

	_, err = l.svcCtx.HelpCategoryService.AddHelpCategory(l.ctx, &cmsclient.AddHelpCategoryReq{
		Name:       req.Name,       // 分类名称
		Icon:       req.Icon,       // 分类图标
		HelpCount:  0,              // 专题数量
		ShowStatus: req.ShowStatus, // 显示状态：0->不显示；1->显示
		Sort:       req.Sort,       // 排序
		CreateBy:   createBy,       // 创建者

	})

	if err != nil {
		logc.Errorf(l.ctx, "添加帮助分类失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.AddHelpCategoryResp{
		Code:    "000000",
		Message: "添加帮助分类成功",
	}, nil
}
