package productbrandservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteProductBrandLogic 删除商品品牌
/*
Author: LiuFeiHua
Date: 2025/05/26 10:33:54
*/
type DeleteProductBrandLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteProductBrandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductBrandLogic {
	return &DeleteProductBrandLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteProductBrand 删除商品品牌
func (l *DeleteProductBrandLogic) DeleteProductBrand(in *pmsclient.DeleteProductBrandReq) (*pmsclient.DeleteProductBrandResp, error) {
	q := query.PmsProductBrand

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除商品品牌失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除商品品牌失败")
	}

	return &pmsclient.DeleteProductBrandResp{}, nil
}
