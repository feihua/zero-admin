package logic

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberAddressAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberAddressAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberAddressAddLogic {
	return MemberAddressAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberAddressAddLogic) MemberAddressAdd(req types.AddMemberAddressReq) (*types.AddMemberAddressResp, error) {
	_, err := l.svcCtx.MemberReceiveAddressService.MemberReceiveAddressAdd(l.ctx, &umsclient.MemberReceiveAddressAddReq{
		MemberId:      req.MemberId,
		Name:          req.Name,
		PhoneNumber:   req.PhoneNumber,
		DefaultStatus: req.DefaultStatus,
		PostCode:      req.PostCode,
		Province:      req.Province,
		City:          req.City,
		Region:        req.Region,
		DetailAddress: req.DetailAddress,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("添加会员地址信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("添加会员地址失败")
	}

	return &types.AddMemberAddressResp{
		Code:    "000000",
		Message: "添加会员地址成功",
	}, nil
}
