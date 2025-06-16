package product_spu

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

// UpdateVerifyStatusLogic 修改审核状态
/*
Author: LiuFeiHua
Date: 2025/6/17 10:42
*/
type UpdateVerifyStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateVerifyStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateVerifyStatusLogic {
	return &UpdateVerifyStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateVerifyStatus 修改审核状态
func (l *UpdateVerifyStatusLogic) UpdateVerifyStatus(req *types.UpdateProductSpuStatusReq) (resp *types.BaseResp, err error) {
	userName := l.ctx.Value("userName").(string)
	_, err = l.svcCtx.ProductSpuService.UpdateVerifyStatus(l.ctx, &pmsclient.UpdateProductSpuStatusReq{
		Ids:       req.Ids,
		Status:    req.Status,
		ReviewMan: userName,
	})

	if err != nil {
		logc.Errorf(l.ctx, "修改审核状态失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return res.Success()
}
