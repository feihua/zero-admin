package subjectcategoryservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type QuerySubjectCategoryDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQuerySubjectCategoryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySubjectCategoryDetailLogic {
	return &QuerySubjectCategoryDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询专题分类表详情
func (l *QuerySubjectCategoryDetailLogic) QuerySubjectCategoryDetail(in *cmsclient.QuerySubjectCategoryDetailReq) (*cmsclient.QuerySubjectCategoryDetailResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.QuerySubjectCategoryDetailResp{}, nil
}
