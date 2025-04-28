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

// AddSubjectCategoryLogic 添加专题分类表
/*
Author: 刘飞华
Date: 2025/02/07 10:21:09
*/
type AddSubjectCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddSubjectCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSubjectCategoryLogic {
	return &AddSubjectCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddSubjectCategory 添加专题分类表
func (l *AddSubjectCategoryLogic) AddSubjectCategory(req *types.AddSubjectCategoryReq) (resp *types.BaseResp, err error) {

	_, err = l.svcCtx.SubjectCategoryService.AddSubjectCategory(l.ctx, &cmsclient.AddSubjectCategoryReq{
		Name:       req.Name,       // 专题分类名称
		Icon:       req.Icon,       // 分类图标
		ShowStatus: req.ShowStatus, // 显示状态：0->不显示；1->显示
		Sort:       req.Sort,       // 排序
		CreateBy:   l.ctx.Value("userName").(string),
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加专题分类表失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return res.Success()
}
