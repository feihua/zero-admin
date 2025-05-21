package address

import (
	"context"
	"github.com/feihua/zero-admin/api/web/internal/logic/common"
	"github.com/feihua/zero-admin/api/web/internal/svc"
	"github.com/feihua/zero-admin/api/web/internal/types"
	"github.com/feihua/zero-admin/pkg/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryMemberAddressDetailLogic 查询会员收货地址详情
/*
Author: LiuFeiHua
Date: 2025/05/21 10:37:06
*/
type QueryMemberAddressDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryMemberAddressDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMemberAddressDetailLogic {
	return &QueryMemberAddressDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryMemberAddressDetail 查询会员收货地址详情
func (l *QueryMemberAddressDetailLogic) QueryMemberAddressDetail(req *types.QueryMemberAddressDetailReq) (resp *types.QueryMemberAddressDetailResp, err error) {
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

	data := types.QueryMemberAddressDetailData{
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

	}
	return &types.QueryMemberAddressDetailResp{
		Code:    0,
		Message: "查询会员收货地址成功",
		Data:    data,
	}, nil
}
