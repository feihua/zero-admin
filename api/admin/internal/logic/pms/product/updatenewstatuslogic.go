package product

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/common/res"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateNewStatusLogic 批量设为新品
/*
Author: LiuFeiHua
Date: 2025/2/5 14:32
*/
type UpdateNewStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateNewStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateNewStatusLogic {
	return &UpdateNewStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateNewStatus 批量设为新品
func (l *UpdateNewStatusLogic) UpdateNewStatus(req *types.UpdateProductStatusReq) (resp *types.BaseResp, err error) {
	_, err = l.svcCtx.ProductService.UpdateNewStatus(l.ctx, &pmsclient.UpdateProductStatusReq{
		Ids:    req.Ids,
		Status: req.Status,
	})

	if err != nil {
		logc.Errorf(l.ctx, "批量设为新品失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("批量设为新品失败")
	}

	return res.Success()
}
