package history

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

// ClearReadHistoryLogic 清空浏览记录
type ClearReadHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewClearReadHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClearReadHistoryLogic {
	return &ClearReadHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ClearReadHistory 清空浏览记录
func (l *ClearReadHistoryLogic) ClearReadHistory() (resp *types.ReadHistoryResp, err error) {
	memberId, err := common.GetMemberId(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.MemberReadHistoryService.DeleteMemberReadHistory(l.ctx, &umsclient.DeleteMemberReadHistoryReq{
		Ids:      nil,
		MemberId: memberId,
	})

	if err != nil {
		logc.Errorf(l.ctx, "清空浏览记录失败,参数memberId: %+v,异常：%s", memberId, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.ReadHistoryResp{
		Code:    0,
		Message: "清空浏览记录成功",
	}, nil
}
