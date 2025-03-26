package memberlevelservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMemberLevelStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMemberLevelStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemberLevelStatusLogic {
	return &UpdateMemberLevelStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新会员等级状态
func (l *UpdateMemberLevelStatusLogic) UpdateMemberLevelStatus(in *umsclient.UpdateMemberLevelStatusReq) (*umsclient.UpdateMemberLevelStatusResp, error) {
	// todo: add your logic here and delete this line

	return &umsclient.UpdateMemberLevelStatusResp{}, nil
}
