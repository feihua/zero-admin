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

// AddCommentReplayLogic 添加产品评价回复表
/*
Author: LiuFeiHua
Date: 2024/6/12 16:33
*/
type AddCommentReplayLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCommentReplayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCommentReplayLogic {
	return &AddCommentReplayLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddCommentReplay 添加产品评价回复表
func (l *AddCommentReplayLogic) AddCommentReplay(in *pmsclient.AddCommentReplayReq) (*pmsclient.AddCommentReplayResp, error) {
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

	return &pmsclient.AddCommentReplayResp{}, nil
}
