package brandservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// BrandDeleteLogic 商品品牌
/*
Author: LiuFeiHua
Date: 2024/5/8 9:47
*/
type BrandDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBrandDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BrandDeleteLogic {
	return &BrandDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// BrandDelete 删除商品品牌
func (l *BrandDeleteLogic) BrandDelete(in *pmsclient.BrandDeleteReq) (*pmsclient.BrandDeleteResp, error) {
	q := query.PmsBrand
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		return nil, err
	}

	return &pmsclient.BrandDeleteResp{}, nil
}
