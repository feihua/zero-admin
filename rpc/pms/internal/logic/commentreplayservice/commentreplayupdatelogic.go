package commentreplayservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"time"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// CommentReplayUpdateLogic 评价回复
/*
Author: LiuFeiHua
Date: 2024/5/8 10:41
*/
type CommentReplayUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentReplayUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentReplayUpdateLogic {
	return &CommentReplayUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CommentReplayUpdate 更新评价回复
func (l *CommentReplayUpdateLogic) CommentReplayUpdate(in *pmsclient.CommentReplayUpdateReq) (*pmsclient.CommentReplayUpdateResp, error) {
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

	return &pmsclient.CommentReplayUpdateResp{}, nil
}
