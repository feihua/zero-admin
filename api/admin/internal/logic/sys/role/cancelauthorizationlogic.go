package role

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/common/res"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// CancelAuthorizationLogic 取消授权/确认授权
/*
Author: LiuFeiHua
Date: 2024/6/02 17:59
*/
type CancelAuthorizationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCancelAuthorizationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelAuthorizationLogic {
	return &CancelAuthorizationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// CancelAuthorization 取消授权/确认授权
func (l *CancelAuthorizationLogic) CancelAuthorization(req *types.CancelAuthorizationReq) (resp *types.BaseResp, err error) {
	_, err = l.svcCtx.RoleService.CancelAuthorization(l.ctx, &sysclient.CancelAuthorizationReq{
		RoleId:  req.RoleId,  // 角色id
		UserIds: req.UserIds, // 用户ids
		IsExist: req.IsExist, // 授权标志
	})

	if err != nil {
		logc.Errorf(l.ctx, "取消授权/确认授权失败,参数:%+v,异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return res.Success()
}
