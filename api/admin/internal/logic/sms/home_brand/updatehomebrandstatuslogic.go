package home_brand

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

// UpdateHomeBrandStatusLogic 首页品牌信息
/*
Author: LiuFeiHua
Date: 2024/5/13 15:54
*/
type UpdateHomeBrandStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateHomeBrandStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateHomeBrandStatusLogic {
	return &UpdateHomeBrandStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateHomeBrandStatus 批量修改推荐品牌状态
// 1.修改sms_home_bran的记录推荐状态(sms-rpc)
// 2.修改cms_brand记录的推荐状态(cms-rpc)
func (l *UpdateHomeBrandStatusLogic) UpdateHomeBrandStatus(req *types.UpdateHomeBrandStatusReq) (resp *types.BaseResp, err error) {
	_, err = l.svcCtx.HomeBrandService.UpdateHomeBrandStatus(l.ctx, &smsclient.UpdateHomeBrandStatusReq{
		Ids:             req.Ids,
		RecommendStatus: req.RecommendStatus, // 推荐状态：0->不推荐;1->推荐
	})

	if err != nil {
		logc.Errorf(l.ctx, "批量修改推荐品牌状态失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	// 2.修改品牌的推荐状态(pms-rpc)
	_, err = l.svcCtx.ProductBrandService.UpdateBrandRecommendStatus(l.ctx, &pmsclient.UpdateProductBrandStatusReq{
		Ids:    req.BrandIds,
		Status: req.RecommendStatus,
	})
	if err != nil {
		logc.Errorf(l.ctx, "根据Ids: %+v,修改品牌的推荐状态异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("修改首页品牌状态失败")
	}
	return res.Success()
}
