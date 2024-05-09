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

// CommentReplayAddLogic 评价回复
/*
Author: LiuFeiHua
Date: 2024/5/8 10:39
*/
type CommentReplayAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentReplayAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentReplayAddLogic {
	return &CommentReplayAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CommentReplayAdd 添加评价回复
func (l *CommentReplayAddLogic) CommentReplayAdd(in *pmsclient.CommentReplayAddReq) (*pmsclient.CommentReplayAddResp, error) {
	err := query.PmsCommentReplay.WithContext(l.ctx).Create(&model.PmsCommentReplay{
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

	return &pmsclient.CommentReplayAddResp{}, nil
}
