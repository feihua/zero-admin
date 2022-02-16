package logic

import (
	"context"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

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
	_, err := l.svcCtx.Ums.MemberReceiveAddressDelete(l.ctx, &umsclient.MemberReceiveAddressDeleteReq{
		Id: req.Id,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据Id: %d,删除会员地址异常:%s", req.Id, err.Error())
		return nil, errorx.NewDefaultError("删除会员地址失败")
	}
	return &types.DeleteMemberAddressResp{
		Code:    "000000",
		Message: "删除会员地址成功",
	}, nil
}
