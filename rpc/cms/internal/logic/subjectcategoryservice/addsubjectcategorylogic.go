package subjectcategoryservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSubjectCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddSubjectCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSubjectCategoryLogic {
	return &AddSubjectCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加专题分类表
func (l *AddSubjectCategoryLogic) AddSubjectCategory(in *cmsclient.AddSubjectCategoryReq) (*cmsclient.AddSubjectCategoryResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.AddSubjectCategoryResp{}, nil
}
