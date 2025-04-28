package subject_category

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/common/res"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteSubjectCategoryLogic 删除专题分类表
/*
Author: 刘飞华
Date: 2025/02/07 10:21:09
*/
type DeleteSubjectCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSubjectCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSubjectCategoryLogic {
	return &DeleteSubjectCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteSubjectCategory 删除专题分类表
func (l *DeleteSubjectCategoryLogic) DeleteSubjectCategory(req *types.DeleteSubjectCategoryReq) (resp *types.BaseResp, err error) {
	_, err = l.svcCtx.SubjectCategoryService.DeleteSubjectCategory(l.ctx, &cmsclient.DeleteSubjectCategoryReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "删除专题分类表失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return res.Success()
}
