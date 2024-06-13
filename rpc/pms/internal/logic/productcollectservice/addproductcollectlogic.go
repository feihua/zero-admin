package productcollectservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddProductCollectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddProductCollectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProductCollectLogic {
	return &AddProductCollectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加收藏表
func (l *AddProductCollectLogic) AddProductCollect(in *pmsclient.AddProductCollectReq) (*pmsclient.AddProductCollectResp, error) {
	// todo: add your logic here and delete this line

	return &pmsclient.AddProductCollectResp{}, nil
}
