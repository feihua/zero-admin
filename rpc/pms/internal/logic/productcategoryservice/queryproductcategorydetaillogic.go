package productcategoryservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryProductCategoryDetailLogic 查询产品分类详情
/*
Author: LiuFeiHua
Date: 2024/6/12 16:56
*/
type QueryProductCategoryDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryProductCategoryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryProductCategoryDetailLogic {
	return &QueryProductCategoryDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryProductCategoryDetail 查询产品分类详情
func (l *QueryProductCategoryDetailLogic) QueryProductCategoryDetail(in *pmsclient.QueryProductCategoryDetailReq) (*pmsclient.QueryProductCategoryDetailResp, error) {

	return &pmsclient.QueryProductCategoryDetailResp{}, nil
}
