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

// UpdateMemberAddressStatusLogic 更新会员收货地址
/*
Author: LiuFeiHua
Date: 2025/05/21 10:37:06
*/
type UpdateMemberAddressStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMemberAddressStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemberAddressStatusLogic {
	return &UpdateMemberAddressStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateMemberAddressStatus 更新会员收货地址状态
func (l *UpdateMemberAddressStatusLogic) UpdateMemberAddressStatus(in *umsclient.UpdateMemberAddressStatusReq) (*umsclient.UpdateMemberAddressStatusResp, error) {
	q := query.UmsMemberAddress

	_, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.Id), q.MemberID.Eq(in.MemberId)).Update(q.IsDefault, in.IsDefault)

	if err != nil {
		logc.Errorf(l.ctx, "更新会员收货地址状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新会员收货地址状态失败")
	}

	return &umsclient.UpdateMemberAddressStatusResp{}, nil
}
