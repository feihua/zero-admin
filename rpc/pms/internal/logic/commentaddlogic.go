package logic

import (
	"context"
	"go-zero-admin/rpc/model/pmsmodel"
	"time"

	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type CommentAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentAddLogic {
	return &CommentAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommentAddLogic) CommentAdd(in *pms.CommentAddReq) (*pms.CommentAddResp, error) {
	CreateTime, _ := time.Parse("2006-01-02 15:04:05", in.CreateTime)
	_, err := l.svcCtx.PmsCommentModel.Insert(pmsmodel.PmsComment{
		ProductId:        in.ProductId,
		MemberNickName:   in.MemberNickName,
		ProductName:      in.ProductName,
		Star:             in.Star,
		MemberIp:         in.MemberIp,
		CreateTime:       CreateTime,
		ShowStatus:       in.ShowStatus,
		ProductAttribute: in.ProductAttribute,
		CollectCouont:    in.CollectCouont,
		ReadCount:        in.ReadCount,
		Content:          in.Content,
		Pics:             in.Pics,
		MemberIcon:       in.MemberIcon,
		ReplayCount:      in.ReplayCount,
	})
	if err != nil {
		return nil, err
	}

	return &pms.CommentAddResp{}, nil
}
