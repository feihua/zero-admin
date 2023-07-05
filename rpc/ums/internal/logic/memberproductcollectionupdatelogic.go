package logic

import (
	"context"

	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberProductCollectionUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberProductCollectionUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberProductCollectionUpdateLogic {
	return &MemberProductCollectionUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberProductCollectionUpdateLogic) MemberProductCollectionUpdate(in *umsclient.MemberProductCollectionUpdateReq) (*umsclient.MemberProductCollectionUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &umsclient.MemberProductCollectionUpdateResp{}, nil
}
