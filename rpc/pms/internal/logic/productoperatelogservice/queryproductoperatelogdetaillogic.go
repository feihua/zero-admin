package productoperatelogservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryProductOperateLogDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryProductOperateLogDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductOperateLogDetailLogic {
	return &QueryProductOperateLogDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询详情
func (l *QueryProductOperateLogDetailLogic) QueryProductOperateLogDetail(in *pmsclient.QueryProductOperateLogDetailReq) (*pmsclient.QueryProductOperateLogDetailResp, error) {
	// todo: add your logic here and delete this line

	return &pmsclient.QueryProductOperateLogDetailResp{}, nil
}
