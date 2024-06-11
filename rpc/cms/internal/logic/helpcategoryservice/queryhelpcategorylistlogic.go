package helpcategoryservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryHelpCategoryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryHelpCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryHelpCategoryListLogic {
	return &QueryHelpCategoryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询帮助分类表列表
func (l *QueryHelpCategoryListLogic) QueryHelpCategoryList(in *cmsclient.QueryHelpCategoryListReq) (*cmsclient.QueryHelpCategoryListResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.QueryHelpCategoryListResp{}, nil
}
