package home_brand

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/common/res"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteHomeBrandLogic 首页品牌信息
/*
Author: LiuFeiHua
Date: 2024/5/13 15:53
*/
type DeleteHomeBrandLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteHomeBrandLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteHomeBrandLogic {
	return DeleteHomeBrandLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteHomeBrand 删除首页品牌信息
func (l *DeleteHomeBrandLogic) DeleteHomeBrand(req *types.DeleteHomeBrandReq) (*types.BaseResp, error) {
	_, err := l.svcCtx.ProductBrandService.UpdateBrandRecommendStatus(l.ctx, &pmsclient.UpdateProductBrandStatusReq{
		Ids:    req.BrandIds,
		Status: 0, // 推荐状态：0->不推荐;1->推荐
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据Ids: %+v,修改品牌的推荐状态异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return res.Success()
}
