package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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
	_, err := l.svcCtx.MemberPriceService.MemberPriceAdd(l.ctx, &pmsclient.MemberPriceAddReq{
		ProductId:       req.ProductId,
		MemberLevelId:   req.MemberLevelId,
		MemberPrice:     req.MemberPrice,
		MemberLevelName: req.MemberLevelName,
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加会员价格信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("添加会员价格失败")
	}

	return &types.AddMemberPriceResp{
		Code:    "000000",
		Message: "添加会员价格成功",
	}, nil
}
