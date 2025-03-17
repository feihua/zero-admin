package subject_category

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

// UpdateSubjectCategoryStatusLogic 更新专题分类表状态状态
/*
Author: 刘飞华
Date: 2025/02/07 10:21:09
*/
type UpdateSubjectCategoryStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSubjectCategoryStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSubjectCategoryStatusLogic {
	return &UpdateSubjectCategoryStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateSubjectCategoryStatus 更新专题分类表状态
func (l *UpdateSubjectCategoryStatusLogic) UpdateSubjectCategoryStatus(req *types.UpdateSubjectCategoryStatusReq) (resp *types.UpdateSubjectCategoryStatusResp, err error) {
	_, err = l.svcCtx.SubjectCategoryService.UpdateSubjectCategoryStatus(l.ctx, &cmsclient.UpdateSubjectCategoryStatusReq{
		Ids:        req.Ids,        // 主键ID
		ShowStatus: req.ShowStatus, // 显示状态：0->不显示；1->显示
		UpdateBy:   l.ctx.Value("userName").(string),
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新专题分类表状态失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.UpdateSubjectCategoryStatusResp{
		Code:    "000000",
		Message: "更新专题分类表状态成功",
	}, nil
}
