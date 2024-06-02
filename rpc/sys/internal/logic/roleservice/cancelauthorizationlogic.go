package roleservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// CancelAuthorizationLogic 取消授权
/*
Author: LiuFeiHua
Date: 2024/6/02 18:31
*/
type CancelAuthorizationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCancelAuthorizationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CancelAuthorizationLogic {
	return &CancelAuthorizationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CancelAuthorization 取消授权
func (l *CancelAuthorizationLogic) CancelAuthorization(in *sysclient.CancelAuthorizationReq) (*sysclient.CancelAuthorizationResp, error) {

	userRole := query.SysUserRole
	_, err := userRole.WithContext(l.ctx).Where(userRole.RoleID.Eq(in.RoleId), userRole.UserID.Eq(in.UserId)).Delete()
	if err != nil {
		logc.Errorf(l.ctx, "取消授权失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("取消授权失败")
	}

	return &sysclient.CancelAuthorizationResp{}, nil
}
