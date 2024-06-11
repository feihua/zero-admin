package subjectcategoryservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSubjectCategoryStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSubjectCategoryStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSubjectCategoryStatusLogic {
	return &UpdateSubjectCategoryStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新专题分类表状态
func (l *UpdateSubjectCategoryStatusLogic) UpdateSubjectCategoryStatus(in *cmsclient.UpdateSubjectCategoryStatusReq) (*cmsclient.UpdateSubjectCategoryStatusResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.UpdateSubjectCategoryStatusResp{}, nil
}
