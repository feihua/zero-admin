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

// DeleteAttentionLogic //取消品牌关注
/*
Author: LiuFeiHua
Date: 2024/5/16 11:03
*/
type DeleteAttentionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteAttentionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAttentionLogic {
	return &DeleteAttentionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteAttention //取消品牌关注
func (l *DeleteAttentionLogic) DeleteAttention(req *types.DeleteAttentionReq) (resp *types.DeleteAttentionResp, err error) {
	memberId, err := common.GetMemberId(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.MemberBrandAttentionService.DeleteMemberBrandAttention(l.ctx, &umsclient.DeleteMemberBrandAttentionReq{
		Ids:      req.BrandIds,
		MemberId: memberId,
	})

	if err != nil {
		logc.Errorf(l.ctx, "删除会员关注的品牌失败,参数: %+v,异常：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.DeleteAttentionResp{
		Code:    0,
		Message: "操作成功",
	}, nil
}
