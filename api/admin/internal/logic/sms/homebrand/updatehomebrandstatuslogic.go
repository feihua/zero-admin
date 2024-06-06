package homebrand

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
func (l *UpdateHomeBrandStatusLogic) UpdateHomeBrandStatus(req *types.UpdateHomeBrandStatusReq) (resp *types.UpdateHomeBrandStatusResp, err error) {
	_, err = l.svcCtx.HomeBrandService.UpdateHomeBrandStatus(l.ctx, &smsclient.UpdateHomeBrandStatusReq{
		Ids:             req.Ids,
		RecommendStatus: req.RecommendStatus,
	})

	if err != nil {
		logc.Errorf(l.ctx, "批量修改推荐品牌状态失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("批量修改推荐品牌状态失败")
	}

	// 2.修改品牌的推荐状态(pms-rpc)
	_, err = l.svcCtx.BrandService.UpdateBrandRecommendStatus(l.ctx, &pmsclient.UpdateBrandRecommendStatusReq{
		Ids:             req.BrandIds,
		RecommendStatus: req.RecommendStatus,
	})
	if err != nil {
		logc.Errorf(l.ctx, "根据Ids: %+v,修改品牌的推荐状态异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("修改首页品牌状态失败")
	}
	return &types.UpdateHomeBrandStatusResp{
		Code:    "000000",
		Message: "批量修改推荐品牌状态成功",
	}, nil
}
