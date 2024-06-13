package feighttemplateservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryFeightTemplateDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryFeightTemplateDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryFeightTemplateDetailLogic {
	return &QueryFeightTemplateDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询运费模版详情
func (l *QueryFeightTemplateDetailLogic) QueryFeightTemplateDetail(in *pmsclient.QueryFeightTemplateDetailReq) (*pmsclient.QueryFeightTemplateDetailResp, error) {
	// todo: add your logic here and delete this line

	return &pmsclient.QueryFeightTemplateDetailResp{}, nil
}
