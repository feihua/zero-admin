package membertagservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMemberTagStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMemberTagStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemberTagStatusLogic {
	return &UpdateMemberTagStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新用户标签表状态
func (l *UpdateMemberTagStatusLogic) UpdateMemberTagStatus(in *umsclient.UpdateMemberTagStatusReq) (*umsclient.UpdateMemberTagStatusResp, error) {
	// todo: add your logic here and delete this line

	return &umsclient.UpdateMemberTagStatusResp{}, nil
}
