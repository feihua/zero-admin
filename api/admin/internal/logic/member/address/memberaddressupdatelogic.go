package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberAddressUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberAddressUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberAddressUpdateLogic {
	return MemberAddressUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberAddressUpdateLogic) MemberAddressUpdate(req types.UpdateMemberAddressReq) (*types.UpdateMemberAddressResp, error) {
	_, err := l.svcCtx.MemberReceiveAddressService.MemberReceiveAddressUpdate(l.ctx, &umsclient.MemberReceiveAddressUpdateReq{
		Id:            req.Id,
		MemberId:      req.MemberId,
		MemberName:    req.Name,
		PhoneNumber:   req.PhoneNumber,
		DefaultStatus: req.DefaultStatus,
		PostCode:      req.PostCode,
		Province:      req.Province,
		City:          req.City,
		Region:        req.Region,
		DetailAddress: req.DetailAddress,
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新会员地址信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("删除会员地址失败")
	}

	return &types.UpdateMemberAddressResp{
		Code:    "000000",
		Message: "删除会员地址成功",
	}, nil
}
