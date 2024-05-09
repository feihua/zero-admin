package logic

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

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
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("添加会员价格信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("添加会员价格失败")
	}

	return &types.AddMemberPriceResp{
		Code:    "000000",
		Message: "添加会员价格成功",
	}, nil
}
