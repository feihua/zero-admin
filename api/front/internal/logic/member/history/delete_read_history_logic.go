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

// DeleteReadHistoryLogic 删除浏览记录
/*
Author: LiuFeiHua
Date: 2025/6/19 11:01
*/
type DeleteReadHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteReadHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteReadHistoryLogic {
	return &DeleteReadHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteReadHistory 删除浏览记录
func (l *DeleteReadHistoryLogic) DeleteReadHistory(req *types.ReadHistoryDeleteReq) (resp *types.ReadHistoryResp, err error) {
	memberId, err := common.GetMemberId(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.MemberReadHistoryService.DeleteMemberReadHistory(l.ctx, &umsclient.DeleteMemberReadHistoryReq{
		Ids:      req.Ids,
		MemberId: memberId,
	})

	if err != nil {
		logc.Errorf(l.ctx, "删除浏览记录失败,参数: %+v,异常：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.ReadHistoryResp{
		Code:    0,
		Message: "删除浏览记录",
	}, nil
}
