package logic

import (
	"context"
	"go-zero-admin/rpc/pms/pmsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MemberPriceAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberPriceAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberPriceAddLogic {
	return MemberPriceAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberPriceAddLogic) MemberPriceAdd(req types.AddMemberPriceReq) (*types.AddMemberPriceResp, error) {
	_, err := l.svcCtx.Pms.MemberPriceAdd(l.ctx, &pmsclient.MemberPriceAddReq{
		ProductId:       req.ProductId,
		MemberLevelId:   req.MemberLevelId,
		MemberPrice:     int64(req.MemberPrice),
		MemberLevelName: req.MemberLevelName,
	})

	if err != nil {
		return nil, err
	}

	return &types.AddMemberPriceResp{
		Code:    "000000",
		Message: "",
	}, nil
}
