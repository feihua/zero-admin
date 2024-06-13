package productcollectservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryProductCollectListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryProductCollectListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductCollectListLogic {
	return &QueryProductCollectListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询收藏表列表
func (l *QueryProductCollectListLogic) QueryProductCollectList(in *pmsclient.QueryProductCollectListReq) (*pmsclient.QueryProductCollectListResp, error) {
	// todo: add your logic here and delete this line

	return &pmsclient.QueryProductCollectListResp{}, nil
}
