package loginlogservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryLoginLogDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryLoginLogDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryLoginLogDetailLogic {
	return &QueryLoginLogDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询系统登录日志表详情
func (l *QueryLoginLogDetailLogic) QueryLoginLogDetail(in *sysclient.QueryLoginLogDetailReq) (*sysclient.QueryLoginLogDetailResp, error) {
	// todo: add your logic here and delete this line

	return &sysclient.QueryLoginLogDetailResp{}, nil
}
