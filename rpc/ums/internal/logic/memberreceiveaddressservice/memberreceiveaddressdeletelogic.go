package memberreceiveaddressservicelogic

import (
	"context"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/rpc/ums/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberReceiveAddressDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberReceiveAddressDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberReceiveAddressDeleteLogic {
	return &MemberReceiveAddressDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberReceiveAddressDeleteLogic) MemberReceiveAddressDelete(in *umsclient.MemberReceiveAddressDeleteReq) (*umsclient.MemberReceiveAddressDeleteResp, error) {
	err := l.svcCtx.UmsMemberReceiveAddressModel.DeleteByIdsAndMemberId(l.ctx, in.Ids, in.MemberId)

	if err != nil {
		return nil, err
	}

	return &umsclient.MemberReceiveAddressDeleteResp{}, nil
}
