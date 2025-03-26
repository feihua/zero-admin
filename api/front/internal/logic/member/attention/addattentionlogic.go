package attention

import (
	"context"
	"encoding/json"
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
Date: 2024/5/16 11:03
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
func (l *AddAttentionLogic) AddAttention(req *types.AddAttentionReq) (resp *types.AddAttentionResp, err error) {
	memberId, _ := l.ctx.Value("memberId").(json.Number).Int64()
	_, err = l.svcCtx.MemberBrandAttentionService.AddMemberBrandAttention(l.ctx, &umsclient.AddMemberBrandAttentionReq{
		BrandId:   req.BrandId,
		BrandName: req.BrandName,
		BrandLogo: req.BrandLogo,
		BrandCity: req.BrandCity,
		MemberId:  memberId,
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加会员关注的品牌失败,参数: %+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.AddAttentionResp{
		Code:    0,
		Message: "操作成功",
	}, nil
}
