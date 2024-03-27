package memberreadhistoryservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberReadHistoryUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberReadHistoryUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberReadHistoryUpdateLogic {
	return &MemberReadHistoryUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberReadHistoryUpdateLogic) MemberReadHistoryUpdate(in *umsclient.MemberReadHistoryUpdateReq) (*umsclient.MemberReadHistoryUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &umsclient.MemberReadHistoryUpdateResp{}, nil
}
