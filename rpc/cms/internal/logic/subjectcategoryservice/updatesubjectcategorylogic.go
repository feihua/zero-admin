package subjectcategoryservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSubjectCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSubjectCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSubjectCategoryLogic {
	return &UpdateSubjectCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新专题分类表
func (l *UpdateSubjectCategoryLogic) UpdateSubjectCategory(in *cmsclient.UpdateSubjectCategoryReq) (*cmsclient.UpdateSubjectCategoryResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.UpdateSubjectCategoryResp{}, nil
}
