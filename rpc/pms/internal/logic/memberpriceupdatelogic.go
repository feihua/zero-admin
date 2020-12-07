package logic

import (
	"context"
	"go-zero-admin/rpc/model/pmsmodel"

	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberPriceUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMemberPriceUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MemberPriceUpdateLogic {
	return &MemberPriceUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MemberPriceUpdateLogic) MemberPriceUpdate(in *pms.MemberPriceUpdateReq) (*pms.MemberPriceUpdateResp, error) {
	err := l.svcCtx.PmsMemberPriceModel.Update(pmsmodel.PmsMemberPrice{
		Id:              in.Id,
		ProductId:       in.ProductId,
		MemberLevelId:   in.MemberLevelId,
		MemberPrice:     float64(in.MemberPrice),
		MemberLevelName: in.MemberLevelName,
	})
	if err != nil {
		return nil, err
	}

	return &pms.MemberPriceUpdateResp{}, nil
}
