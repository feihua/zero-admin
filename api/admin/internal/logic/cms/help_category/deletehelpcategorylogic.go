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

// DeleteHelpCategoryLogic 删除帮助分类
/*
Author: 刘飞华
Date: 2026/08/27 10:50:29
*/
type DeleteHelpCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteHelpCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteHelpCategoryLogic {
	return &DeleteHelpCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteHelpCategory 删除帮助分类
func (l *DeleteHelpCategoryLogic) DeleteHelpCategory(req *types.DeleteHelpCategoryReq) (resp *types.DeleteHelpCategoryResp, err error) {
	_, err = l.svcCtx.HelpCategoryService.DeleteHelpCategory(l.ctx, &cmsclient.DeleteHelpCategoryReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "删除帮助分类失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.DeleteHelpCategoryResp{
		Code:    "000000",
		Message: "删除帮助分类成功",
	}, nil
}
