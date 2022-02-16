package logic

import (
	"context"
	"encoding/json"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/pms/pmsclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

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
	_, err := l.svcCtx.Pms.MemberPriceAdd(l.ctx, &pmsclient.MemberPriceAddReq{
		ProductId:       req.ProductId,
		MemberLevelId:   req.MemberLevelId,
		MemberPrice:     int64(req.MemberPrice),
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
