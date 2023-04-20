package address

import (
	"context"
	"encoding/json"
	"zero-admin/common/ctxdata"
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

func (l *AddressDeleteLogic) AddressDelete(req *types.AddressDeleteReq) (resp *types.AddressDeleteResp, err error) {
	memberId := ctxdata.GetUidFromCtx(l.ctx)
	addressQueryDetail, err := l.svcCtx.Ums.MemberReceiveAddressQueryDetail(l.ctx, &umsclient.MemberReceiveAddressQueryDetailReq{
		MemberId:  memberId,
		AddressID: req.AddressID,
	})

	if addressQueryDetail == nil || err != nil {
		// reqStr, _ := json.Marshal(req)
		// logx.WithContext(l.ctx).Errorf("查询会员收货详细地址失败,参数：%s,响应：%s", reqStr, err.Error())
		return &types.AddressDeleteResp{
			Code: 1,
			Msg:  "参数有误",
		}, nil
	}
	_, err = l.svcCtx.Ums.MemberReceiveAddressDelete(l.ctx, &umsclient.MemberReceiveAddressDeleteReq{
		Id: req.AddressID,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("删除会员收货地址失败,参数：%s,响应：%s", reqStr, err.Error())
		return &types.AddressDeleteResp{
			Code: 1,
			Msg:  err.Error(),
		}, nil
	}

	return &types.AddressDeleteResp{
		Code: 0,
		Msg:  "删除会员收货地址成功",
	}, nil

}
