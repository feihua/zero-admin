package helpservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryHelpDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryHelpDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryHelpDetailLogic {
	return &QueryHelpDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询帮助表详情
func (l *QueryHelpDetailLogic) QueryHelpDetail(in *cmsclient.QueryHelpDetailReq) (*cmsclient.QueryHelpDetailResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.QueryHelpDetailResp{}, nil
}
