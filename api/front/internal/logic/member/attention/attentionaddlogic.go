package attention

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// AttentionAddLogic
/*
Author: LiuFeiHua
Date: 2023/12/4 17:10
*/
type AttentionAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAttentionAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AttentionAddLogic {
	return &AttentionAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AttentionAdd 添加会员关注信息
func (l *AttentionAddLogic) AttentionAdd(req *types.AddAttentionReq) (resp *types.AddAttentionResp, err error) {
	memberId, _ := l.ctx.Value("memberId").(json.Number).Int64()
	_, err = l.svcCtx.MemberAttentionService.MemberBrandAttentionAdd(l.ctx, &umsclient.MemberBrandAttentionAddReq{
		BrandId:   req.BrandId,
		BrandName: req.BrandName,
		BrandLogo: req.BrandLogo,
		BrandCity: req.BrandCity,
		MemberId:  memberId,
	})
	if err != nil {
		return nil, err
	}

	return &types.AddAttentionResp{
		Code:    0,
		Message: "操作成功",
	}, nil
}
