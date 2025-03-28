package memberservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

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
	q := query.UmsMember
	_, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.Id)).Update(q.Integration, in.Integration)
	if err != nil {
		logc.Errorf(l.ctx, "更新会员积分失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新会员积分失败")
	}

	return &umsclient.UpdateMemberIntegrationResp{}, nil
}
