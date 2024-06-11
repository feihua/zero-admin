package helpcategoryservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryHelpCategoryDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryHelpCategoryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryHelpCategoryDetailLogic {
	return &QueryHelpCategoryDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询帮助分类表详情
func (l *QueryHelpCategoryDetailLogic) QueryHelpCategoryDetail(in *cmsclient.QueryHelpCategoryDetailReq) (*cmsclient.QueryHelpCategoryDetailResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.QueryHelpCategoryDetailResp{}, nil
}
