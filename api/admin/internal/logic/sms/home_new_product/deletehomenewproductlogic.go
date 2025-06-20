package home_new_product

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

func NewDeleteHomeNewProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomeNewProductDeleteLogic {
	return HomeNewProductDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteHomeNewProduct 删除首页新品
func (l *HomeNewProductDeleteLogic) DeleteHomeNewProduct(req *types.DeleteHomeNewProductReq) (*types.BaseResp, error) {

	_, err := l.svcCtx.ProductSpuService.UpdateNewStatus(l.ctx, &pmsclient.UpdateProductSpuStatusReq{
		Ids:    req.ProductIds,
		Status: 0, // 推荐状态：0->不推荐;1->推荐
	})
	if err != nil {
		logc.Errorf(l.ctx, "根据Ids: %+v,修改新鲜好物状态异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return res.Success()
}
