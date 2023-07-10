package memberreadhistoryservicelogic

import (
	"context"

	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberReadHistoryDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberReadHistoryDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberReadHistoryDeleteLogic {
	return &MemberReadHistoryDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberReadHistoryDeleteLogic) MemberReadHistoryDelete(in *umsclient.MemberReadHistoryDeleteReq) (*umsclient.MemberReadHistoryDeleteResp, error) {
	err := l.svcCtx.UmsMemberReadHistoryModel.DeleteByIdsAndMemberId(l.ctx, in.Ids, in.MemberId)

	if err != nil {
		return nil, err
	}

	return &umsclient.MemberReadHistoryDeleteResp{}, nil
}
