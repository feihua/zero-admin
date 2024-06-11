package subjectcategoryservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type QuerySubjectCategoryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQuerySubjectCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySubjectCategoryListLogic {
	return &QuerySubjectCategoryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询专题分类表列表
func (l *QuerySubjectCategoryListLogic) QuerySubjectCategoryList(in *cmsclient.QuerySubjectCategoryListReq) (*cmsclient.QuerySubjectCategoryListResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.QuerySubjectCategoryListResp{}, nil
}
