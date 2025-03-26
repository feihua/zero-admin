package memberreceiveaddressservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteMemberReceiveAddressLogic 删除会员收货地址
/*
Author: LiuFeiHua
Date: 2024/6/11 14:04
*/
type DeleteMemberReceiveAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMemberReceiveAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMemberReceiveAddressLogic {
	return &DeleteMemberReceiveAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteMemberReceiveAddress 删除会员收货地址
func (l *DeleteMemberReceiveAddressLogic) DeleteMemberReceiveAddress(in *umsclient.DeleteMemberReceiveAddressReq) (*umsclient.DeleteMemberReceiveAddressResp, error) {
	q := query.UmsMemberReceiveAddress
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...), q.MemberID.Eq(in.MemberId)).Delete()
	if err != nil {
		logc.Errorf(l.ctx, "删除会员收货地址失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除会员收货地址失败")
	}

	return &umsclient.DeleteMemberReceiveAddressResp{}, nil
}
