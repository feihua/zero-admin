package memberreceiveaddressservicelogic

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/rpc/ums/internal/svc"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// MemberReceiveAddressListLogic
/*
Author: LiuFeiHua
Date: 2023/11/29 16:17
*/
type MemberReceiveAddressListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberReceiveAddressListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberReceiveAddressListLogic {
	return &MemberReceiveAddressListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// MemberReceiveAddressList 查询会员收货地址列表
func (l *MemberReceiveAddressListLogic) MemberReceiveAddressList(in *umsclient.MemberReceiveAddressListReq) (*umsclient.MemberReceiveAddressListResp, error) {
	all, err := l.svcCtx.UmsMemberReceiveAddressModel.FindAll(l.ctx, in)
	count, _ := l.svcCtx.UmsMemberReceiveAddressModel.Count(l.ctx, in)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logc.Errorf(l.ctx, "查询会员地址列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}
	var list []*umsclient.MemberReceiveAddressListData
	for _, item := range *all {

		list = append(list, &umsclient.MemberReceiveAddressListData{
			Id:            item.Id,
			MemberId:      item.MemberId,
			Name:          item.Name,
			PhoneNumber:   item.PhoneNumber,
			DefaultStatus: item.DefaultStatus,
			PostCode:      item.PostCode,
			Province:      item.Province,
			City:          item.City,
			Region:        item.Region,
			DetailAddress: item.DetailAddress,
		})
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(list)
	logc.Infof(l.ctx, "查询会员地址列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &umsclient.MemberReceiveAddressListResp{
		Total: count,
		List:  list,
	}, nil

}
