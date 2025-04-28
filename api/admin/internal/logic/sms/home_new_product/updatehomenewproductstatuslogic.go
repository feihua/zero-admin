package home_new_product

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/common/res"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

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

func NewUpdateHomeNewProductStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateNewProductStatusLogic {
	return &UpdateNewProductStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateHomeNewProductStatus 批量修改推荐状态
// 1.修改sms_home_new_product的记录(sms-rpc)
// 2.修改pms_product记录的状态(pms-rpc)
func (l *UpdateNewProductStatusLogic) UpdateHomeNewProductStatus(req *types.UpdateHomeNewProductStatusReq) (resp *types.BaseResp, err error) {
	// 1.修改sms_home_new_product的记录(sms-rpc)
	_, err = l.svcCtx.HomeNewProductService.UpdateHomeNewProductStatus(l.ctx, &smsclient.UpdateHomeNewProductStatusReq{
		Ids:             req.Ids,
		RecommendStatus: req.RecommendStatus, // 推荐状态：0->不推荐;1->推荐
	})

	if err != nil {
		logc.Errorf(l.ctx, "批量修改推荐状态失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	// 2.修改pms_product记录的状态(pms-rpc)
	_, err = l.svcCtx.ProductService.UpdateNewStatus(l.ctx, &pmsclient.UpdateProductStatusReq{
		Ids:    req.ProductIds,
		Status: req.RecommendStatus,
	})
	if err != nil {
		logc.Errorf(l.ctx, "根据Ids: %+v,修改新鲜好物状态异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}
	return res.Success()

}
