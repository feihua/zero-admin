package memberservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMemberStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMemberStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemberStatusLogic {
	return &UpdateMemberStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新会员表状态
func (l *UpdateMemberStatusLogic) UpdateMemberStatus(in *umsclient.UpdateMemberStatusReq) (*umsclient.UpdateMemberStatusResp, error) {
	// todo: add your logic here and delete this line

	return &umsclient.UpdateMemberStatusResp{}, nil
}
