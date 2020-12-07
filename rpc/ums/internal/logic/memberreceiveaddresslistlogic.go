package logic

import (
	"context"
	"fmt"

	"go-zero-admin/rpc/ums/internal/svc"
	"go-zero-admin/rpc/ums/ums"

	"github.com/tal-tech/go-zero/core/logx"
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
	all, _ := l.svcCtx.UmsMemberReceiveAddressModel.FindAll(in.Current, in.PageSize)
	//count, _ := l.svcCtx.UserModel.Count()

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

	fmt.Println(list)
	return &ums.MemberReceiveAddressListResp{
		Total: 10,
		List:  list,
	}, nil

}
