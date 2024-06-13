package productvertifyrecordservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryProductVertifyRecordDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryProductVertifyRecordDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductVertifyRecordDetailLogic {
	return &QueryProductVertifyRecordDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询商品审核记录详情
func (l *QueryProductVertifyRecordDetailLogic) QueryProductVertifyRecordDetail(in *pmsclient.QueryProductVertifyRecordDetailReq) (*pmsclient.QueryProductVertifyRecordDetailResp, error) {
	// todo: add your logic here and delete this line

	return &pmsclient.QueryProductVertifyRecordDetailResp{}, nil
}
