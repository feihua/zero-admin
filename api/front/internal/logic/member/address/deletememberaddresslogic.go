package address

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteMemberAddressLogic 删除收货地址
/*
Author: LiuFeiHua
Date: 2024/5/16 13:59
*/
type DeleteMemberAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteMemberAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMemberAddressLogic {
	return &DeleteMemberAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteMemberAddress 删除收货地址
func (l *DeleteMemberAddressLogic) DeleteMemberAddress(req *types.DeleteMemberAddressReq) (resp *types.DeleteMemberAddressResp, err error) {
	memberId, _ := l.ctx.Value("memberId").(json.Number).Int64()
	_, _ = l.svcCtx.MemberReceiveAddressService.MemberReceiveAddressDelete(l.ctx, &umsclient.MemberReceiveAddressDeleteReq{
		Id:       req.Id,
		MemberId: memberId,
	})

	return &types.DeleteMemberAddressResp{
		Code:    0,
		Message: "操作成功",
	}, nil
}
