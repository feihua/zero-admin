package address

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/feihua/zero-admin/front-api/internal/svc"
	"github.com/feihua/zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// MemberAddressDeleteLogic 收货地址相关
/*
Author: LiuFeiHua
Date: 2023/11/30 11:21
*/
type MemberAddressDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberAddressDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberAddressDeleteLogic {
	return &MemberAddressDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// MemberAddressDelete 删除会员收货地址
func (l *MemberAddressDeleteLogic) MemberAddressDelete(req *types.DeleteMemberAddressReq) (resp *types.DeleteMemberAddressResp, err error) {
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
