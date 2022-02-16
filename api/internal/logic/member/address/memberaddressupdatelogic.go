package logic

import (
	"context"
	"encoding/json"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/ums/umsclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

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
	_, err := l.svcCtx.Ums.MemberReceiveAddressUpdate(l.ctx, &umsclient.MemberReceiveAddressUpdateReq{
		Id:            req.Id,
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
		logx.WithContext(l.ctx).Errorf("更新会员地址信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("删除会员地址失败")
	}

	return &types.UpdateMemberAddressResp{
		Code:    "000000",
		Message: "删除会员地址成功",
	}, nil
}
