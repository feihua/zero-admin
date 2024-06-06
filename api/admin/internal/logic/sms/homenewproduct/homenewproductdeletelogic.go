package homenewproduct

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// HomeNewProductDeleteLogic 首页新品
/*
Author: LiuFeiHua
Date: 2024/5/14 9:17
*/
type HomeNewProductDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeNewProductDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomeNewProductDeleteLogic {
	return HomeNewProductDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// HomeNewProductDelete 删除首页新品
// 1.删除sms_home_new_product的记录(sms-rpc)
// 2.修改pms_product记录的状态为不推荐(pms-rpc)
func (l *HomeNewProductDeleteLogic) HomeNewProductDelete(req types.DeleteHomeNewProductReq) (*types.DeleteHomeNewProductResp, error) {
	// 1.删除sms_home_new_product的记录(sms-rpc)
	_, err := l.svcCtx.HomeNewProductService.HomeNewProductDelete(l.ctx, &smsclient.HomeNewProductDeleteReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据Id: %+v,删除新鲜好物异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("删除新鲜好物失败")
	}

	// 2.修改pms_product记录的状态为不推荐(pms-rpc)
	_, err = l.svcCtx.ProductService.UpdateNewStatus(l.ctx, &pmsclient.UpdateProductStatusReq{
		Ids:    req.ProductIds,
		Status: 0,
	})
	if err != nil {
		logc.Errorf(l.ctx, "根据Ids: %+v,修改新鲜好物状态异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("删除新鲜好物失败")
	}

	return &types.DeleteHomeNewProductResp{
		Code:    "000000",
		Message: "删除新鲜好物成功",
	}, nil
}
