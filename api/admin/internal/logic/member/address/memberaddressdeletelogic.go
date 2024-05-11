package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberAddressDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberAddressDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberAddressDeleteLogic {
	return MemberAddressDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberAddressDeleteLogic) MemberAddressDelete(req types.DeleteMemberAddressReq) (*types.DeleteMemberAddressResp, error) {
	_, err := l.svcCtx.MemberReceiveAddressService.MemberReceiveAddressDelete(l.ctx, &umsclient.MemberReceiveAddressDeleteReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据Id: %+v,删除会员地址异常:%s", req.Id, err.Error())
		return nil, errorx.NewDefaultError("删除会员地址失败")
	}
	return &types.DeleteMemberAddressResp{
		Code:    "000000",
		Message: "删除会员地址成功",
	}, nil
}
