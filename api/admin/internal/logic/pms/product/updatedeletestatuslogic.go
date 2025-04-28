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

// UpdateDeleteStatusLogic 批量修改删除状态
/*
Author: LiuFeiHua
Date: 2025/2/5 14:31
*/
type UpdateDeleteStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateDeleteStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDeleteStatusLogic {
	return &UpdateDeleteStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateDeleteStatus 批量修改删除状态
func (l *UpdateDeleteStatusLogic) UpdateDeleteStatus(req *types.UpdateProductStatusReq) (resp *types.BaseResp, err error) {
	_, err = l.svcCtx.ProductService.UpdateDeleteStatus(l.ctx, &pmsclient.UpdateProductStatusReq{
		Ids:    req.Ids,
		Status: req.Status,
	})

	if err != nil {
		logc.Errorf(l.ctx, "批量修改删除状态失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("批量修改删除状态失败")
	}

	return res.Success()
}
