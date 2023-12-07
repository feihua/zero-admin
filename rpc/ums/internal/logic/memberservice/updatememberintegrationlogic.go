package memberservicelogic

import (
	"context"

	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateMemberIntegrationLogic
/*
Author: LiuFeiHua
Date: 2023/12/7 13:54
*/
type UpdateMemberIntegrationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMemberIntegrationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemberIntegrationLogic {
	return &UpdateMemberIntegrationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateMemberIntegration 更新会员积分
func (l *UpdateMemberIntegrationLogic) UpdateMemberIntegration(in *umsclient.UpdateMemberIntegrationReq) (*umsclient.UpdateMemberIntegrationResp, error) {
	err := l.svcCtx.UmsMemberModel.UpdateMemberIntegration(l.ctx, in.Id, in.Integration)

	if err != nil {
		return nil, err
	}

	return &umsclient.UpdateMemberIntegrationResp{}, nil
}
