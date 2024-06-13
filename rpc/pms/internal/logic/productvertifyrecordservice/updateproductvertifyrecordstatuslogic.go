package productvertifyrecordservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProductVertifyRecordStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductVertifyRecordStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductVertifyRecordStatusLogic {
	return &UpdateProductVertifyRecordStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新商品审核记录状态
func (l *UpdateProductVertifyRecordStatusLogic) UpdateProductVertifyRecordStatus(in *pmsclient.UpdateProductVertifyRecordStatusReq) (*pmsclient.UpdateProductVertifyRecordStatusResp, error) {
	// todo: add your logic here and delete this line

	return &pmsclient.UpdateProductVertifyRecordStatusResp{}, nil
}
