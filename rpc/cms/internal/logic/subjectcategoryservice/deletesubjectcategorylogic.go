package subjectcategoryservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSubjectCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteSubjectCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSubjectCategoryLogic {
	return &DeleteSubjectCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除专题分类表
func (l *DeleteSubjectCategoryLogic) DeleteSubjectCategory(in *cmsclient.DeleteSubjectCategoryReq) (*cmsclient.DeleteSubjectCategoryResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.DeleteSubjectCategoryResp{}, nil
}
