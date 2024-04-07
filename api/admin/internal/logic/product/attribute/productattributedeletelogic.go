package attribute

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductAttributeDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductAttributeDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductAttributeDeleteLogic {
	return &ProductAttributeDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductAttributeDeleteLogic) ProductAttributeDelete(req *types.DeleteProductAttributeReq) (resp *types.DeleteProductAttributeResp, err error) {
	_, err = l.svcCtx.ProductAttributeService.ProductAttributeDelete(l.ctx, &pmsclient.ProductAttributeDeleteReq{
		Ids: req.Ids,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据Id: %d,删除属性异常:%s", req.Ids, err.Error())
		return nil, errorx.NewDefaultError("删除属性失败")
	}

	return &types.DeleteProductAttributeResp{
		Code:    "000000",
		Message: "",
	}, nil
}
