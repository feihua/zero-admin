package preferredareaservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryPreferredAreaDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryPreferredAreaDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryPreferredAreaDetailLogic {
	return &QueryPreferredAreaDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询优选专区详情
func (l *QueryPreferredAreaDetailLogic) QueryPreferredAreaDetail(in *cmsclient.QueryPreferredAreaDetailReq) (*cmsclient.QueryPreferredAreaDetailResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.QueryPreferredAreaDetailResp{}, nil
}
