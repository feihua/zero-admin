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

// AddMemberAddressLogic 添加会员收货地址
/*
Author: LiuFeiHua
Date: 2025/05/21 10:37:06
*/
type AddMemberAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddMemberAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMemberAddressLogic {
	return &AddMemberAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddMemberAddress 添加会员收货地址
func (l *AddMemberAddressLogic) AddMemberAddress(req *types.MemberAddressReq) (resp *types.MemberAddressResp, err error) {
	memberId, err := common.GetMemberId(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.MemberAddressService.AddMemberAddress(l.ctx, &umsclient.AddMemberAddressReq{
		MemberId:      memberId,          // 会员ID
		ReceiverName:  req.ReceiverName,  // 收货人姓名
		ReceiverPhone: req.ReceiverPhone, // 收货人电话
		Province:      req.Province,      // 省份
		City:          req.City,          // 城市
		District:      req.District,      // 区县
		DetailAddress: req.DetailAddress, // 详细地址
		PostalCode:    req.PostalCode,    // 邮政编码
		Tag:           req.Tag,           // 地址标签：家、公司等
		IsDefault:     req.IsDefault,     // 是否默认地址
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加会员收货地址失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.MemberAddressResp{
		Code:    0,
		Message: "添加会员收货地址成功",
	}, nil
}
