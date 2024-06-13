package commentreplayservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"time"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateCommentReplayLogic 更新产品评价回复表
/*
Author: LiuFeiHua
Date: 2024/6/12 16:35
*/
type UpdateCommentReplayLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCommentReplayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCommentReplayLogic {
	return &UpdateCommentReplayLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateCommentReplay 更新产品评价回复表
func (l *UpdateCommentReplayLogic) UpdateCommentReplay(in *pmsclient.UpdateCommentReplayReq) (*pmsclient.UpdateCommentReplayResp, error) {
	q := query.PmsCommentReplay
	_, err := q.WithContext(l.ctx).Updates(&model.PmsCommentReplay{
		ID:             in.Id,
		CommentID:      in.CommentId,
		MemberNickName: in.MemberNickName,
		MemberIcon:     in.MemberIcon,
		Content:        in.Content,
		CreateTime:     time.Now(),
		Type:           in.Type,
	})

	if err != nil {
		return nil, err
	}

	return &pmsclient.UpdateCommentReplayResp{}, nil
}
