package address

import (
	"context"
	"encoding/json"
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

	_, err = l.svcCtx.Ums.MemberReceiveAddressDelete(l.ctx, &umsclient.MemberReceiveAddressDeleteReq{
		Ids: req.AddressID,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("删除会员收货地址失败,参数：%s,响应：%s", reqStr, err.Error())
		return &types.AddressDeleteResp{
			Errno:  1,
			Errmsg: err.Error(),
		}, nil
	}

	return &types.AddressDeleteResp{
		Errno:  0,
		Errmsg: "删除会员收货地址成功",
	}, nil

}
