package productcollectservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteProductCollectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteProductCollectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductCollectLogic {
	return &DeleteProductCollectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除收藏表
func (l *DeleteProductCollectLogic) DeleteProductCollect(in *pmsclient.DeleteProductCollectReq) (*pmsclient.DeleteProductCollectResp, error) {
	// todo: add your logic here and delete this line

	return &pmsclient.DeleteProductCollectResp{}, nil
}
