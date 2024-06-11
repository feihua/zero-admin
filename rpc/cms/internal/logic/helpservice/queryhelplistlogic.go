package helpservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryHelpListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryHelpListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryHelpListLogic {
	return &QueryHelpListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询帮助表列表
func (l *QueryHelpListLogic) QueryHelpList(in *cmsclient.QueryHelpListReq) (*cmsclient.QueryHelpListResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.QueryHelpListResp{}, nil
}
