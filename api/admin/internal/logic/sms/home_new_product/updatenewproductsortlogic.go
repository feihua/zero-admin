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

// UpdateNewProductSortLogic
/*
Author: LiuFeiHua
Date: 2024/5/14 9:18
*/
type UpdateNewProductSortLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateNewProductSortLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateNewProductSortLogic {
	return &UpdateNewProductSortLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateNewProductSort 修改推荐排序
func (l *UpdateNewProductSortLogic) UpdateNewProductSort(req *types.UpdateNewProductSortReq) (resp *types.BaseResp, err error) {
	_, err = l.svcCtx.ProductService.UpdateProductSort(l.ctx, &pmsclient.UpdateProductSortReq{
		Id:   req.ProductId,
		Sort: req.Sort,
	})

	if err != nil {
		logc.Errorf(l.ctx, "修改推荐排序失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return res.Success()
}
