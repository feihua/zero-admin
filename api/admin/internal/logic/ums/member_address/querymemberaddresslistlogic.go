package member_address

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
	result, err := l.svcCtx.MemberAddressService.QueryMemberAddressList(l.ctx, &umsclient.QueryMemberAddressListReq{
		PageNum:  req.Current,
		PageSize: req.PageSize,
		MemberId: req.MemberId,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询会员地址列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询会员地址失败")
	}

	var list []*types.ListMemberAddressData

	for _, detail := range result.List {
		list = append(list, &types.ListMemberAddressData{
			Id:            detail.Id,            // 主键ID
			MemberId:      detail.MemberId,      // 会员ID
			ReceiverName:  detail.ReceiverName,  // 收货人姓名
			ReceiverPhone: detail.ReceiverPhone, // 收货人电话
			Province:      detail.Province,      // 省份
			City:          detail.City,          // 城市
			District:      detail.District,      // 区县
			DetailAddress: detail.DetailAddress, // 详细地址
			PostalCode:    detail.PostalCode,    // 邮政编码
			Tag:           detail.Tag,           // 地址标签：家、公司等
			IsDefault:     detail.IsDefault,     // 是否默认地址
			CreateTime:    detail.CreateTime,    // 创建时间
			UpdateTime:    detail.UpdateTime,    // 更新时间
			IsDeleted:     detail.IsDeleted,     // 是否删除
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
