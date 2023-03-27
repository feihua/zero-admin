package logic

import (
	"context"

	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberReceiveAddressQueryDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberReceiveAddressQueryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberReceiveAddressQueryDetailLogic {
	return &MemberReceiveAddressQueryDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberReceiveAddressQueryDetailLogic) MemberReceiveAddressQueryDetail(in *umsclient.MemberReceiveAddressQueryDetailReq) (*umsclient.MemberReceiveAddressQueryDetailResp, error) {
	// todo: add your logic here and delete this line

	return &umsclient.MemberReceiveAddressQueryDetailResp{}, nil
}
