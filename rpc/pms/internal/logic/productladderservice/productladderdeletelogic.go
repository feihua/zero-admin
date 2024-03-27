package productladderservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductLadderDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductLadderDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductLadderDeleteLogic {
	return &ProductLadderDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductLadderDeleteLogic) ProductLadderDelete(in *pmsclient.ProductLadderDeleteReq) (*pmsclient.ProductLadderDeleteResp, error) {
	err := l.svcCtx.PmsProductLadderModel.Delete(l.ctx, in.Id)

	if err != nil {
		return nil, err
	}

	return &pmsclient.ProductLadderDeleteResp{}, nil
}
