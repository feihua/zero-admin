package productvertifyrecordservicelogic

import (
	"context"
	"zero-admin/rpc/pms/pmsclient"

	"zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductVertifyRecordDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductVertifyRecordDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductVertifyRecordDeleteLogic {
	return &ProductVertifyRecordDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductVertifyRecordDeleteLogic) ProductVertifyRecordDelete(in *pmsclient.ProductVertifyRecordDeleteReq) (*pmsclient.ProductVertifyRecordDeleteResp, error) {
	err := l.svcCtx.PmsProductVertifyRecordModel.Delete(l.ctx, in.Id)

	if err != nil {
		return nil, err
	}

	return &pmsclient.ProductVertifyRecordDeleteResp{}, nil
}
