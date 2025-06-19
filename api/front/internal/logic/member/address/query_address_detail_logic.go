package address

import (
	"context"
	"github.com/feihua/zero-admin/api/front/internal/logic/common"
	"github.com/feihua/zero-admin/pkg/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryAddressDetailLogic 查询会员收货地址详情
/*
Author: LiuFeiHua
Date: 2025/6/19 10:49
*/
type QueryAddressDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryAddressDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryAddressDetailLogic {
	return &QueryAddressDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryAddressDetail 查询会员收货地址详情
func (l *QueryAddressDetailLogic) QueryAddressDetail(req *types.QueryAddressDetailReq) (resp *types.QueryAddressDetailResp, err error) {
	memberId, err := common.GetMemberId(l.ctx)
	if err != nil {
		return nil, err
	}
	detail, err := l.svcCtx.MemberAddressService.QueryMemberAddressDetail(l.ctx, &umsclient.QueryMemberAddressDetailReq{
		Id:       req.Id,
		MemberId: memberId,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询会员收货地址详情失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	data := types.QueryAddressDetailData{
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

	}
	return &types.QueryAddressDetailResp{
		Code:    0,
		Message: "查询会员收货地址成功",
		Data:    data,
	}, nil
}
