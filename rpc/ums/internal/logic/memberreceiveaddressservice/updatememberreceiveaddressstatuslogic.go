package memberreceiveaddressservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMemberReceiveAddressStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMemberReceiveAddressStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMemberReceiveAddressStatusLogic {
	return &UpdateMemberReceiveAddressStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新会员收货地址表状态
func (l *UpdateMemberReceiveAddressStatusLogic) UpdateMemberReceiveAddressStatus(in *umsclient.UpdateMemberReceiveAddressStatusReq) (*umsclient.UpdateMemberReceiveAddressStatusResp, error) {
	// todo: add your logic here and delete this line

	return &umsclient.UpdateMemberReceiveAddressStatusResp{}, nil
}
