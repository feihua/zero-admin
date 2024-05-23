package address

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMemberAddressListLogic 会员地址
/*
Author: LiuFeiHua
Date: 2024/5/23 9:09
*/
type QueryMemberAddressListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryMemberAddressListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberAddressListLogic {
	return &QueryMemberAddressListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryMemberAddressList 查询会员地址
func (l *QueryMemberAddressListLogic) QueryMemberAddressList(req *types.ListMemberAddressReq) (resp *types.ListMemberAddressResp, err error) {
	result, err := l.svcCtx.MemberReceiveAddressService.MemberReceiveAddressList(l.ctx, &umsclient.MemberReceiveAddressListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
		MemberId: req.MemberId,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询会员地址列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询会员地址失败")
	}

	var list []*types.ListMemberAddressData

	for _, item := range result.List {
		list = append(list, &types.ListMemberAddressData{
			Id:            item.Id,
			MemberId:      item.MemberId,
			Name:          item.MemberName,
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
		Total:    result.Total,
		Code:     "000000",
		Message:  "查询会员地址成功",
	}, nil
}
