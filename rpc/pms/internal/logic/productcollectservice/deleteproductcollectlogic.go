package productcollectservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteProductCollectLogic 删除收藏
/*
Author: LiuFeiHua
Date: 2025/01/24 09:08:06
*/
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

// DeleteProductCollect 删除收藏
func (l *DeleteProductCollectLogic) DeleteProductCollect(in *pmsclient.DeleteProductCollectReq) (*pmsclient.DeleteProductCollectResp, error) {

	return &pmsclient.DeleteProductCollectResp{}, nil
}
