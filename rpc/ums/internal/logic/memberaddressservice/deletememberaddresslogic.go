package memberaddressservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/ums/gen/query"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteMemberAddressLogic 删除会员收货地址
/*
Author: LiuFeiHua
Date: 2025/05/21 10:37:06
*/
type DeleteMemberAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMemberAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMemberAddressLogic {
	return &DeleteMemberAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteMemberAddress 删除会员收货地址
func (l *DeleteMemberAddressLogic) DeleteMemberAddress(in *umsclient.DeleteMemberAddressReq) (*umsclient.DeleteMemberAddressResp, error) {
	q := query.UmsMemberAddress

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...), q.MemberID.Eq(in.MemberId)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除会员收货地址失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除会员收货地址失败")
	}

	return &umsclient.DeleteMemberAddressResp{}, nil
}
