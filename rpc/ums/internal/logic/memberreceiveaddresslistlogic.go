package logic

import (
	"context"
	"encoding/json"
	"zero-admin/rpc/ums/internal/svc"
	"zero-admin/rpc/ums/ums"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *MemberReceiveAddressListLogic) MemberReceiveAddressList(in *ums.MemberReceiveAddressListReq) (*ums.MemberReceiveAddressListResp, error) {
	all, err := l.svcCtx.UmsMemberReceiveAddressModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.UmsMemberReceiveAddressModel.Count()

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询会员地址列表信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}
	var list []*ums.MemberReceiveAddressListData
	for _, item := range *all {

		list = append(list, &ums.MemberReceiveAddressListData{
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
	logx.WithContext(l.ctx).Infof("查询会员地址列表信息,参数：%s,响应：%s", reqStr, listStr)
	return &ums.MemberReceiveAddressListResp{
		Total: count,
		List:  list,
	}, nil

}
