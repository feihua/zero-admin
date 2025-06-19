package address

import (
	"context"
	"github.com/feihua/zero-admin/api/front/internal/logic/common"
	"github.com/feihua/zero-admin/pkg/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteAddressLogic 删除会员收货地址
/*
Author: LiuFeiHua
Date: 2025/6/19 10:47
*/
type DeleteAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAddressLogic {
	return &DeleteAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteAddress 删除会员收货地址
func (l *DeleteAddressLogic) DeleteAddress(req *types.DeleteAddressReq) (resp *types.AddressResp, err error) {
	memberId, err := common.GetMemberId(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.MemberAddressService.DeleteMemberAddress(l.ctx, &umsclient.DeleteMemberAddressReq{
		Ids:      req.Ids,
		MemberId: memberId,
	})

	if err != nil {
		logc.Errorf(l.ctx, "删除会员收货地址失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.AddressResp{
		Code:    0,
		Message: "删除会员收货地址成功",
	}, nil
}
