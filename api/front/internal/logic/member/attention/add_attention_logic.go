package attention

import (
	"context"
	"github.com/feihua/zero-admin/api/front/internal/logic/common"
	"github.com/feihua/zero-admin/pkg/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddAttentionLogic 添加会员关注的品牌
/*
Author: LiuFeiHua
Date: 2025/6/19 10:56
*/
type AddAttentionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddAttentionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddAttentionLogic {
	return &AddAttentionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddAttention 添加会员关注的品牌
func (l *AddAttentionLogic) AddAttention(req *types.AttentionReq) (resp *types.AttentionResp, err error) {
	memberId, err := common.GetMemberId(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.MemberBrandAttentionService.AddMemberBrandAttention(l.ctx, &umsclient.AddMemberBrandAttentionReq{
		BrandId:   req.BrandId,
		BrandName: req.BrandName,
		BrandLogo: req.BrandLogo,
		BrandCity: req.BrandCity,
		MemberId:  memberId,
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加会员关注的品牌失败,参数: %+v,异常：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.AttentionResp{
		Code:    0,
		Message: "操作成功",
	}, nil
}
