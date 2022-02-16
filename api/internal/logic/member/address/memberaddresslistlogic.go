package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberAddressListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberAddressListLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberAddressListLogic {
	return MemberAddressListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberAddressListLogic) MemberAddressList(req types.ListMemberAddressReq) (*types.ListMemberAddressResp, error) {
	resp, err := l.svcCtx.Ums.MemberReceiveAddressList(l.ctx, &umsclient.MemberReceiveAddressListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	if err != nil {
		data, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("参数: %s,查询会员地址列表异常:%s", string(data), err.Error())
		return nil, errorx.NewDefaultError("查询会员地址失败")
	}

	for _, data := range resp.List {
		fmt.Println(data)
	}
	var list []*types.ListtMemberAddressData

	for _, item := range resp.List {
		list = append(list, &types.ListtMemberAddressData{
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

	return &types.ListMemberAddressResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
		Code:     "000000",
		Message:  "查询会员地址成功",
	}, nil
}
