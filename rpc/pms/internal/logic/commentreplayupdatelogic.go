package logic

import (
	"context"
	"go-zero-admin/rpc/model/pmsmodel"
	"time"

	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *CommentReplayUpdateLogic) CommentReplayUpdate(in *pms.CommentReplayUpdateReq) (*pms.CommentReplayUpdateResp, error) {
	CreateTime, _ := time.Parse("2006-01-02 15:04:05", in.CreateTime)
	err := l.svcCtx.PmsCommentReplayModel.Update(pmsmodel.PmsCommentReplay{
		Id:             in.Id,
		CommentId:      in.CommentId,
		MemberNickName: in.MemberNickName,
		MemberIcon:     in.MemberIcon,
		Content:        in.Content,
		CreateTime:     CreateTime,
		Type:           in.Type,
	})
	if err != nil {
		return nil, err
	}

	return &pms.CommentReplayUpdateResp{}, nil
}
