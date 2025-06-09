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

// HomeNewProductAddLogic 首页新品
/*
Author: LiuFeiHua
Date: 2024/5/14 9:17
*/
type HomeNewProductAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddHomeNewProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomeNewProductAddLogic {
	return HomeNewProductAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddHomeNewProduct 添加首页新品
func (l *HomeNewProductAddLogic) AddHomeNewProduct(req *types.AddHomeNewProductReq) (*types.BaseResp, error) {

	_, err := l.svcCtx.ProductService.UpdateNewStatus(l.ctx, &pmsclient.UpdateProductStatusReq{
		Ids:    req.ProductIds,
		Status: 1,
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据Ids: %+v,添加首页新品异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return res.Success()
}
