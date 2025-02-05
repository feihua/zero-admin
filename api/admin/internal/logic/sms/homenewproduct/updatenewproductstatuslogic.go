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

// UpdateNewProductStatusLogic
/*
Author: LiuFeiHua
Date: 2024/5/14 9:18
*/
type UpdateNewProductStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateNewProductStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateNewProductStatusLogic {
	return &UpdateNewProductStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateNewProductStatus 批量修改推荐状态
// 1.修改sms_home_new_product的记录(sms-rpc)
// 2.修改pms_product记录的状态(pms-rpc)
func (l *UpdateNewProductStatusLogic) UpdateNewProductStatus(req *types.UpdateNewProductStatusReq) (resp *types.UpdateNewProductStatusResp, err error) {
	// 1.修改sms_home_new_product的记录(sms-rpc)
	_, err = l.svcCtx.HomeNewProductService.UpdateHomeNewProductStatus(l.ctx, &smsclient.UpdateHomeNewProductStatusReq{
		Ids:             req.Ids,
		RecommendStatus: req.RecommendStatus, // 推荐状态：0->不推荐;1->推荐
	})

	if err != nil {
		logc.Errorf(l.ctx, "批量修改推荐状态失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("批量修改推荐状态失败")
	}

	// 2.修改pms_product记录的状态(pms-rpc)
	_, err = l.svcCtx.ProductService.UpdateNewStatus(l.ctx, &pmsclient.UpdateProductStatusReq{
		Ids:    req.ProductIds,
		Status: req.RecommendStatus,
	})
	if err != nil {
		logc.Errorf(l.ctx, "根据Ids: %+v,修改新鲜好物状态异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("删除新鲜好物失败")
	}
	return &types.UpdateNewProductStatusResp{
		Code:    "000000",
		Message: "批量修改推荐状态成功",
	}, nil

}
