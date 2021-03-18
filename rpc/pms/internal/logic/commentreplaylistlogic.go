package logic

import (
	"context"
	"fmt"
	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type CommentReplayListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentReplayListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentReplayListLogic {
	return &CommentReplayListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommentReplayListLogic) CommentReplayList(in *pms.CommentReplayListReq) (*pms.CommentReplayListResp, error) {
	all, _ := l.svcCtx.PmsCommentReplayModel.FindAll(in.Current, in.PageSize)
	count, _ := l.svcCtx.PmsCommentReplayModel.Count()

	var list []*pms.CommentReplayListData
	for _, item := range *all {

		list = append(list, &pms.CommentReplayListData{
			Id:             item.Id,
			CommentId:      item.CommentId,
			MemberNickName: item.MemberNickName,
			MemberIcon:     item.MemberIcon,
			Content:        item.Content,
			CreateTime:     item.CreateTime.Format("2006-01-02 15:04:05"),
			Type:           item.Type,
		})
	}

	fmt.Println(list)
	return &pms.CommentReplayListResp{
		Total: count,
		List:  list,
	}, nil
}
