package brandservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryBrandDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryBrandDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryBrandDetailLogic {
	return &QueryBrandDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询品牌表详情
func (l *QueryBrandDetailLogic) QueryBrandDetail(in *pmsclient.QueryBrandDetailReq) (*pmsclient.QueryBrandDetailResp, error) {
	// todo: add your logic here and delete this line

	return &pmsclient.QueryBrandDetailResp{}, nil
}
