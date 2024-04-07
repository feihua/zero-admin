package history

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// ReadHistoryDeleteLogic
/*
Author: LiuFeiHua
Date: 2023/11/30 16:30
*/
type ReadHistoryDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReadHistoryDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReadHistoryDeleteLogic {
	return &ReadHistoryDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ReadHistoryDelete 删除浏览记录
func (l *ReadHistoryDeleteLogic) ReadHistoryDelete(req *types.ReadHistoryDeleteReq) (resp *types.ReadHistoryDeleteResp, err error) {
	memberId, _ := l.ctx.Value("memberId").(json.Number).Int64()
	_, _ = l.svcCtx.MemberReadHistoryService.MemberReadHistoryDelete(l.ctx, &umsclient.MemberReadHistoryDeleteReq{
		Id:       req.Id,
		MemberId: memberId,
	})

	return &types.ReadHistoryDeleteResp{
		Code:    0,
		Message: "操作成功",
	}, nil
}
