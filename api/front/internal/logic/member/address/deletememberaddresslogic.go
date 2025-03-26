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
	memberId, err := common.GetMemberId(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.MemberReceiveAddressService.DeleteMemberReceiveAddress(l.ctx, &umsclient.DeleteMemberReceiveAddressReq{
		Ids:      req.Ids,
		MemberId: memberId,
	})

	if err != nil {
		logc.Errorf(l.ctx, "删除收货地址失败,参数: %+v,异常：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}
	return &types.DeleteMemberAddressResp{
		Code:    0,
		Message: "操作成功",
	}, nil
}
