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

// ClearAttentionLogic 清空当前用户品牌关注列表
/*
Author: LiuFeiHua
Date: 2025/3/26 11:43
*/
type ClearAttentionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewClearAttentionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClearAttentionLogic {
	return &ClearAttentionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ClearAttention 清空当前用户品牌关注列表
func (l *ClearAttentionLogic) ClearAttention() (resp *types.AttentionResp, err error) {
	memberId, err := common.GetMemberId(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.MemberBrandAttentionService.DeleteMemberBrandAttention(l.ctx, &umsclient.DeleteMemberBrandAttentionReq{
		MemberId: memberId,
	})

	if err != nil {
		logc.Errorf(l.ctx, "清空当前用户品牌关注列表失败,参数memberId: %d,异常：%s", memberId, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.AttentionResp{
		Code:    0,
		Message: "操作成功",
	}, nil
}
