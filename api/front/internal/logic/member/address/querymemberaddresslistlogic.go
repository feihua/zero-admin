package address

import (
	"context"
	"github.com/feihua/zero-admin/api/front/internal/logic/common"
	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"
	"github.com/feihua/zero-admin/pkg/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMemberAddressListLogic 查询会员收货地址列表
/*
Author: LiuFeiHua
Date: 2025/05/21 10:37:06
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

// QueryMemberAddressList 查询会员收货地址列表
func (l *QueryMemberAddressListLogic) QueryMemberAddressList(req *types.QueryMemberAddressListReq) (resp *types.QueryMemberAddressListResp, err error) {
	memberId, err := common.GetMemberId(l.ctx)
	if err != nil {
		return nil, err
	}
	result, err := l.svcCtx.MemberAddressService.QueryMemberAddressList(l.ctx, &umsclient.QueryMemberAddressListReq{
		PageNum:  req.Current,
		PageSize: req.PageSize,
		MemberId: memberId, // 会员ID
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询字会员收货地址列表失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.QueryMemberAddressListData

	for _, detail := range result.List {
		list = append(list, &types.QueryMemberAddressListData{
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

	return &types.QueryMemberAddressListResp{
		Code:     0,
		Message:  "查询会员收货地址列表成功",
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    result.Total,
	}, nil
}
