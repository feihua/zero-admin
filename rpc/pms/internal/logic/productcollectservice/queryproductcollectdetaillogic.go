package productcollectservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryProductCollectDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryProductCollectDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductCollectDetailLogic {
	return &QueryProductCollectDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询收藏表详情
func (l *QueryProductCollectDetailLogic) QueryProductCollectDetail(in *pmsclient.QueryProductCollectDetailReq) (*pmsclient.QueryProductCollectDetailResp, error) {
	// todo: add your logic here and delete this line

	return &pmsclient.QueryProductCollectDetailResp{}, nil
}
