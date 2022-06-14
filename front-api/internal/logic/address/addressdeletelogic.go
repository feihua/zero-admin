package address

import (
	"context"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/front-api/internal/svc"
	"zero-admin/front-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddressDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddressDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddressDeleteLogic {
	return AddressDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddressDeleteLogic) AddressDelete(req types.AddressDeleteReq) (resp *types.AddressDeleteResp, err error) {

	_, _ = l.svcCtx.Ums.MemberReceiveAddressDelete(l.ctx, &umsclient.MemberReceiveAddressDeleteReq{
		Id: req.AddressID,
	})

	return &types.AddressDeleteResp{
		Errno:  0,
		Errmsg: "删除会员信息成功",
	}, nil

}
