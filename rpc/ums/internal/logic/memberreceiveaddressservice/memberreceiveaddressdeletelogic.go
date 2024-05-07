package memberreceiveaddressservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// MemberReceiveAddressDeleteLogic 会员收货地址
/*
Author: LiuFeiHua
Date: 2023/11/30 11:24
*/
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

// MemberReceiveAddressDelete 删除会员收货地址
func (l *MemberReceiveAddressDeleteLogic) MemberReceiveAddressDelete(in *umsclient.MemberReceiveAddressDeleteReq) (*umsclient.MemberReceiveAddressDeleteResp, error) {
	q := query.UmsMemberReceiveAddress
	_, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.Id), q.MemberID.Eq(in.MemberId)).Delete()
	if err != nil {
		return nil, err
	}

	return &umsclient.MemberReceiveAddressDeleteResp{}, nil
}
