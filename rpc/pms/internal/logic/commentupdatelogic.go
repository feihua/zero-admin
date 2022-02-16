package logic

import (
	"context"
	"time"
	"zero-admin/rpc/model/pmsmodel"

	"zero-admin/rpc/pms/internal/svc"
	"zero-admin/rpc/pms/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentUpdateLogic {
	return &CommentUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CommentUpdateLogic) CommentUpdate(in *pms.CommentUpdateReq) (*pms.CommentUpdateResp, error) {
	CreateTime, _ := time.Parse("2006-01-02 15:04:05", in.CreateTime)
	err := l.svcCtx.PmsCommentModel.Update(pmsmodel.PmsComment{
		Id:               in.Id,
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

	return &pms.CommentUpdateResp{}, nil
}
