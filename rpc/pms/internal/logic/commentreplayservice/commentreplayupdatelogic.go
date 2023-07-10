package commentreplayservicelogic

import (
	"context"
	"time"
	"zero-admin/rpc/model/pmsmodel"
	"zero-admin/rpc/pms/pmsclient"

	"zero-admin/rpc/pms/internal/svc"

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

func (l *CommentReplayUpdateLogic) CommentReplayUpdate(in *pmsclient.CommentReplayUpdateReq) (*pmsclient.CommentReplayUpdateResp, error) {
	CreateTime, _ := time.Parse("2006-01-02 15:04:05", in.CreateTime)
	err := l.svcCtx.PmsCommentReplayModel.Update(l.ctx, &pmsmodel.PmsCommentReplay{
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

	return &pmsclient.CommentReplayUpdateResp{}, nil
}
